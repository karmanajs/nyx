package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const (
	// Config for Application
	nameApp    = "Nyx"
	nameBin    = "nyx"
	versionApp = "0.0.1"

	// Default parametrs
	defaultPorts    = "80,443"
	defaultProtocol = "tcp"

	// Consts
	maxPort = 65535
)

// Config holds application configuration
type Config struct {
	Host       string
	Ports      string
	Protocol   string
	OutputJSON string
}

// ScanResult represents the result of a port scan
type ScanResult struct {
	Port    int    `json:"port"`
	Status  string `json:"status"`
	Service string `json:"service,omitempty"`
}

func main() {

	config := parseFlags()

	parsedPorts, err := ParsePorts(config.Ports)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing ports: %v\n", err)
		os.Exit(1)
	}

	if config.Host == "" {
		fmt.Fprintf(os.Stdout, "Usage: %s --help for check examples\n", nameBin)
		os.Exit(0)
	}

	if _, err := net.LookupHost(config.Host); err != nil {
		fmt.Fprintf(os.Stderr, "Error: invalid host %s - %v\n", config.Host, err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stdout, "Starting %s %s\n", nameApp, versionApp)
	fmt.Fprintf(os.Stdout, "Scanning %s (%s)\n", config.Host, config.Protocol)

	results := Scan(config.Host, parsedPorts, config.Protocol)

	if config.OutputJSON == "" {
		for _, checkedPort := range results {
			fmt.Fprintf(os.Stdout, "  %d - %s\n", checkedPort.Port, checkedPort.Status)
		}
	} else {
		err := SaveToJSONFile(config.OutputJSON, results)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error saving JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Results saved to %s\n", config.OutputJSON)
	}

	fmt.Fprintf(os.Stdout, "Finish\n")
}

// TODO doc
func parseFlags() Config {
	var config Config

	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, "%s %s\n", nameApp, versionApp)
		fmt.Fprintf(os.Stdout, "A simple port scanner.\n")
		fmt.Fprintf(os.Stdout, "\nUsage:\n")
		fmt.Fprintf(os.Stdout, "    %s -h <host> [-p <ports>] [-tp <protocol>]\n", nameBin)

		fmt.Fprintf(os.Stdout, "\nOptions:\n")
		fmt.Fprintf(os.Stdout, "  -h, --host: Ip address or url to host.\n")
		fmt.Fprintf(os.Stdout, "    Can use ip, ip4 (IPv4-only), ip6 (IPv6-only), url.\n")
		fmt.Fprintf(os.Stdout, "  -p, --ports <port range>: Only scan specified ports.\n")
		fmt.Fprintf(os.Stdout, "  -tp, --type-protocol: Type of internet protocol.\n")
		fmt.Fprintf(os.Stdout, "    All types: tcp, tcp4 (IPv4-only), tcp6 (IPv6-only), udp, udp4 (IPv4-only), udp6 (IPv6-only).\n")
		fmt.Fprintf(os.Stdout, "  -jf, --json-file: Output results in JSON format\n")

		fmt.Fprintf(os.Stdout, "\nExamples:\n")
		fmt.Fprintf(os.Stdout, "  %s -h example.com -p 443,54,70-80\n", nameBin)
		fmt.Fprintf(os.Stdout, "  %s -h=example.com --ports 443,54,70-80\n", nameBin)
		fmt.Fprintf(os.Stdout, "  %s --host=example.com -p=443,54,70-80 -jf out.json\n", nameBin)
		fmt.Fprintf(os.Stdout, "  %s --host=example.com --ports=443,54,70-80 --json-file out.json\n", nameBin)
	}

	flag.StringVar(&config.Host, "h", "", "Target host (IP or domain)")
	flag.StringVar(&config.Host, "host", "", "Target host (IP or domain)")
	flag.StringVar(&config.Ports, "p", defaultPorts, "Ports to scan (comma-separated or range)")
	flag.StringVar(&config.Ports, "ports", defaultPorts, "Ports to scan")
	flag.StringVar(&config.Protocol, "tp", defaultProtocol, "Protocol type (tcp/udp/etc.)")
	flag.StringVar(&config.Protocol, "type-protocol", defaultProtocol, "Protocol type")
	flag.StringVar(&config.OutputJSON, "jf", "", "Output results in JSON format")
	flag.StringVar(&config.OutputJSON, "json-file", "", "Output results in JSON format")

	flag.Parse()

	return config
}

// TODO doc
func Scan(host string, ports []int, protocol string) []ScanResult {

	var (
		results []ScanResult
		mu      sync.Mutex
		wg      sync.WaitGroup
	)

	for _, pr := range ports {
		wg.Add(1)
		go func(goPort int) {
			defer wg.Done()
			socket := fmt.Sprintf(host+":%d", goPort)
			conn, err := net.Dial(protocol, socket)

			result := ScanResult{Port: goPort}

			if err != nil {
				result.Status = "closed"
			} else {
				result.Status = "open"
				conn.Close()
			}

			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}(pr)
	}

	wg.Wait()
	return results
}

// ParsePorts: TODO doc
func ParsePorts(inputPorts string) ([]int, error) {

	// check for invalids characters
	if matched, _ := regexp.MatchString(`[^0-9,\-]`, inputPorts); matched {
		return nil, fmt.Errorf("invalid characters in port specification")
	}

	var ports []int
	portsRanges := strings.Split(inputPorts, ",")

	for _, pr := range portsRanges {

		if pr == "" {
			continue
		}

		if strings.Contains(pr, "-") {
			rangePorts, err := ParseRangePorts(pr)
			if err != nil {
				return nil, err
			}

			ports = append(ports, rangePorts...)
		} else {
			port, err := ParseSinglePort(pr)
			if err != nil {
				return nil, err
			}

			ports = append(ports, port)
		}

	}

	return DeduplicatePorts(ports), nil
}

// TODO doc
func ParseRangePorts(pr string) ([]int, error) {
	// handle port range
	parts := strings.Split(pr, "-")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invaliv port range format: %s", pr)
	}
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid start port in range %s: %v", pr, err)
	}
	if start < 0 || start > maxPort {
		return nil, fmt.Errorf("start port out of range (0-%d): %d", maxPort, start)
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid end port in range %s: %v", pr, err)
	}
	if end < 0 || end > maxPort {
		return nil, fmt.Errorf("end port out of range (0-%d): %d", maxPort, end)
	}
	if start > end {
		return nil, fmt.Errorf("start port cannot be greater than end port in range %s", pr)
	}

	var ports []int
	for port := start; port <= end; port++ {
		ports = append(ports, port)
	}

	return ports, nil
}

// TODO doc
func ParseSinglePort(pr string) (int, error) {
	port, err := strconv.Atoi(pr)
	if err != nil {
		return 0, fmt.Errorf("invalid port number: %s", pr)
	}
	if port < 0 || port > maxPort {
		return 0, fmt.Errorf("port out of range (0-%d): %d", maxPort, port)
	}

	return port, nil
}

// TODO doc
func DeduplicatePorts(ports []int) []int {
	keys := make(map[int]bool)
	var uniquePorts []int

	for _, port := range ports {
		if _, value := keys[port]; !value {
			keys[port] = true
			uniquePorts = append(uniquePorts, port)
		}
	}

	return uniquePorts
}

// TODO doc
func SaveToJSONFile(filename string, data []ScanResult) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(data); err != nil {
		return fmt.Errorf("failed to encode JSON: %w", err)
	}

	return nil
}
