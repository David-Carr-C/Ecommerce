import { Button } from '@mui/material';
import './App.css'
import { useState } from 'react';
import { Checkbox } from '@nextui-org/checkbox';

function App() {
  const [count, setCount] = useState(0);

  return (
    <>
      <p>Hello World!</p>
      <Button variant="contained" onClick={() => setCount(count + 1)} sx={buttonClass}>
        Hello World {count}
      </Button>
      <div className="text-center">
        <img src="https://picsum.photos/200/300" alt="random" className='rounded-full inline-block' />
      </div>
      Checkbox from NextUI
      <Checkbox defaultChecked color='danger' className='m-2 uppercase'>
        Checkbox
      </Checkbox>
    </>
  );
}

export default App

const buttonClass = {
  backgroundColor: 'white',
  color: 'gray',
  '&:hover': {
    backgroundColor: 'gray'
  }
}