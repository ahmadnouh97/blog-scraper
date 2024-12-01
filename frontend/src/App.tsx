import React from 'react'
import Layout from './app/layout'


const App = () => {
  const content = <div>Hello World</div> 
  return (
    <>
    <Layout children={content} />
    </>
  )
}

export default App
