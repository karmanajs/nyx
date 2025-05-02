import React, { useState } from 'react';
import './App.css';

export default function App() {
  const [form, setForm] = useState({
    host: '',
    ports: '',
    protocol: 'tcp',
    timeout: '2s'
  });
  
  const [results, setResults] = useState([]);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    
    try {
      const response = await fetch('http://localhost:8080/scan', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(form)
      });
      
      if (!response.ok) throw new Error('Scan failed');
      setResults(await response.json());
      
    } catch (error) {
      alert(error.message);
      console.error('Scan error:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm(prev => ({ ...prev, [name]: value }));
  };

  // Вынесенные компоненты формы и результатов
  const ScanForm = () => (
    <form className="scanner-form" onSubmit={handleSubmit}>
      <FormField
        id="host"
        label="TARGET HOST"
        value={form.host}
        onChange={handleChange}
        placeholder="example.com or 192.168.1.1"
      />
      
      <FormField
        id="ports"
        label="PORTS RANGE"
        value={form.ports}
        onChange={handleChange}
        placeholder="80,443,8080-8090"
      />
      
      <div className="form-group">
        <label htmlFor="protocol">PROTOCOL</label>
        <select
          id="protocol"
          name="protocol"
          value={form.protocol}
          onChange={handleChange}
        >
          <option value="tcp">TCP</option>
          <option value="udp">UDP</option>
          <option value="tcp4">TCP (IPv4 ONLY)</option>
          <option value="udp6">UDP (IPv6 ONLY)</option>
        </select>
      </div>
      
      <FormField
        id="timeout"
        label="TIMEOUT"
        value={form.timeout}
        onChange={handleChange}
        placeholder="2s / 400ms"
      />
      
      <button 
        type="submit" 
        className="scan-button"
        disabled={loading}
      >
        {loading ? 'SCANNING...' : 'INITIATE SCAN'}
      </button>
    </form>
  );

  const ResultsView = () => (
    <div className="results-container">
      <h2>Scan Results</h2>
      <table className="results-table">
        <thead>
          <tr><th>Port</th><th>Status</th></tr>
        </thead>
        <tbody>
          {results.map((item, i) => (
            <tr key={i}>
              <td>{item.port}</td>
              <td className={item.status.toLowerCase()}>{item.status}</td>
            </tr>
          ))}
        </tbody>
      </table>
      <button 
        onClick={() => setResults([])} 
        className="back-button"
      >
        New Scan
      </button>
    </div>
  );

  return (
    <div className="App">
      <header className="header">
        <h1>NYX PORT SCANNER</h1>
      </header>

      <main className="main">
        {results.length > 0 ? <ResultsView /> : <ScanForm />}
      </main>

      <footer className="footer">
        NYX NETWORK SECURITY TOOL | v0.0.1
      </footer>
    </div>
  );
}

// Дополнительный компонент для полей ввода
function FormField({ id, label, value, onChange, placeholder }) {
  return (
    <div className="form-group">
      <label htmlFor={id}>{label}</label>
      <input
        type="text"
        id={id}
        name={id}
        value={value}
        onChange={onChange}
        placeholder={placeholder}
        required
      />
    </div>
  );
}