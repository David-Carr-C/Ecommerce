import './App.css'
import { Routes, Route, Link } from 'react-router-dom';
import Home from './pages/Home';
import { Divider, Drawer, ListItem, ListItemButton, ListItemIcon, ListItemText } from '@mui/material';
import ProofParam from './pages/ProofParam';
import NotFound from './pages/NotFound';

function App() {
  return (
    <>
      <Drawer variant="permanent" anchor="right">
        <Link to="/">
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
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/param/:id" element={<ProofParam />} />
        <Route path="*" element={<NotFound />} />
      </Routes>
    </>
  );
}

export default App