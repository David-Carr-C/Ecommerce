import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'
import './index.css'
import 'tailwindcss/tailwind.css'
import { NextUIProvider } from '@nextui-org/react'
import { BrowserRouter } from 'react-router-dom'

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode> {/* El modo estricto es una herramienta para detectar errores potenciales en una aplicación de React. */}
    <BrowserRouter> {/* El componente BrowserRouter es un componente de enrutamiento que envuelve nuestra aplicación y nos permite usar las funciones de enrutamiento de React. */}
      <NextUIProvider> {/* El componente NextUIProvider es un componente de NextUI que envuelve nuestra aplicación y nos permite usar las funciones de NextUI. */}
        <App /> {/* El componente App es nuestra aplicación. */}
      </NextUIProvider>
    </BrowserRouter>
  </React.StrictMode>
);
