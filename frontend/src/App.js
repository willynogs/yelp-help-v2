import React, { useState } from 'react'

const Test = () => {
  const [active, setActive] = useState(false)
  return (
    <>
      <p>{active ? ':: ACTIVE ::' : 'Inactive'}</p>
      <button type='button' onClick={() => setActive(prevState => !prevState)}>Toggle</button>
    </>
  )
}

export default Test
