import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'
import './index.css'
import MyRoutes from './routes'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
    <React.StrictMode>
        <MyRoutes />
    </React.StrictMode>
)
