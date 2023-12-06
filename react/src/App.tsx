import { Button } from '@mui/material';
import './App.css'
import { useState } from 'react';

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <p>Hello World!</p>
      <Button variant="contained" onClick={() => setCount(count + 1)}>
        Hello World {count}
      </Button>
    </>
  );
}

export default App
