import React from 'react'
import styled from 'styled-components'
import useAxios from 'axios-hooks'

const PageTitle = styled.h1`
  font-size: 4rem;
  text-align: center;
`

const Button = styled.button`
  margin: 0 auto;
  display: block;
  height: 35px;
  border: 1px solid black;
  font-size: 0.9rem;
  border-radius: 6px;
`

const Home = () => {
  const [{ data, loading, err }, execute] = useAxios({
    url: 'http://localhost:8000/random',
    method: 'GET'
  }, {
    manual: true
  })

  console.log(data, loading, err)

  return (
    <div>
      <PageTitle>Yelp Help v2</PageTitle>
      <Button type='button' onClick={() => execute()}>Find a Restaurant</Button>
    </div>
  )
}

export default Home
