import React from 'react'
import styled from 'styled-components'
import Router from './components/Router'

const Container = styled.div`
  background: rgb(124,16,16);
  background: linear-gradient(135deg, rgba(124,16,16,1) 0%, rgba(211,35,35,1) 100%);
  color: #ffffff;
  font-family: 'Mukta', sans-serif;
  position: absolute;
  min-height: 100vh;
  min-width: 100%;
`

const App = () => {
  return (
    <Container>
      <Router />
    </Container>
  )
}

export default App
