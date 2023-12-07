import { Button } from '@mui/material';
import { useState } from 'react';
import { Checkbox } from '@nextui-org/checkbox';
import { Avatar } from '@nextui-org/avatar';
import { Link } from 'react-router-dom';

export default function Home() {
  const [count, setCount] = useState(0)

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
      <Avatar
        src="https://picsum.photos/200/300"
        alt="random"
        className='m-2'
      />
      <Link to="/404">Go to other</Link>
    </>
  )
}

const buttonClass = {
  backgroundColor: 'white',
  color: 'gray',
  '&:hover': {
    backgroundColor: 'gray'
  }
}