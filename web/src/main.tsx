import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import App from './components/App.tsx'
import './index.css'
import { BrowserRouter as Router } from 'react-router-dom'
import axios from 'axios'

axios.defaults.baseURL = 'http://localhost:4000/'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Router>
      <App />
    </Router>
  </StrictMode>,
)
