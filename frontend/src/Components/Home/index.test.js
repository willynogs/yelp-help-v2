import '@testing-library/jest-dom'
import React from 'react'
import { render, fireEvent, screen } from '@testing-library/react'
import Home from '../Home'

test('it renders the apps title', () => {
  render(<Home />)
  expect(screen.getByText('Yelp Help v2')).toBeInTheDocument()
})

test('it renders a button', () => {
  render(<Home />)
  expect(screen.getByText('Find a Restaurant')).toBeInTheDocument()
})
