import './App.css'
import { Routes, Route, Link } from 'react-router-dom';
import Home from './pages/Home';
import { Divider, Drawer, ListItem, ListItemButton, ListItemIcon, ListItemText } from '@mui/material';
import ProofParam from './pages/ProofParam';
import NotFound from './pages/NotFound';

function App() {
  return (
    <>
      <Drawer variant="permanent" anchor="right"> {/* Drawer es un menu lateral */}
        <Link to="/"> {/* Link es un componente que permite navegar entre paginas */}
          <ListItem>
            <ListItemButton>
              <ListItemIcon>
                <img
                  src="https://avatars.githubusercontent.com/u/63006416?v=4"
                  alt="avatar"
                  width="50"
                  height="50"
                />
              </ListItemIcon>
              <ListItemText primary="Dylan" />
            </ListItemButton>
          </ListItem>
        </Link>
        <Divider />
        <Link to="/param/23">
          <ListItem>
            <ListItemButton>
              <ListItemIcon>
                <img
                  src="https://avatars.githubusercontent.com/u/63006416?v=4"
                  alt="avatar"
                  width="50"
                  height="50"
                />
              </ListItemIcon>
              <ListItemText primary="Dylan" />
            </ListItemButton>
          </ListItem>
        </Link>
      </Drawer>
      <Routes> {/* Aqui van las rutas, y que componente debe renderizar */}
        <Route path="/" element={<Home />} />
        <Route path="/param/:id" element={<ProofParam />} /> {/* :id es un parametro */}
        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
}

export default App