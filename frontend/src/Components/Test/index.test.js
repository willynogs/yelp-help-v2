import '@testing-library/jest-dom'
import React from 'react'
import { render, fireEvent, screen } from '@testing-library/react'
import Test from '../Test'

test('this is just a test', () => {
  render(<Test />)
  expect(screen.queryByText(':: ACTIVE ::')).toBeNull()
  fireEvent.click(screen.getByText('Toggle'))
  expect(screen.getByText(':: ACTIVE ::')).toBeInTheDocument()
})
