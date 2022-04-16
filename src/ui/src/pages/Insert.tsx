import React from 'react';
import { useState } from 'react'
import Button from "@mui/material/Button";

function Header({title} : {title:any}) {
  return <h1>{title ? title : 'Default title'}</h1>
}


const Insert = () => {
  const names = ['Ada Lovelace', 'Grace Hopper', 'Margaret Hamilton']

  const [likes, setLikes] = useState(0)

  function handleClick() {
    setLikes(likes + 1)
  }

  return (
    <div>
      <Header title="Develop. Preview. Ship. 🚀" />
      <ul>
        {names.map(name => (
          <li key={name}>{name}</li>
        ))}
      </ul>

      <Button variant="contained" onClick={handleClick}>Like ({likes})</Button>
    </div>
  )
}


export default Insert;