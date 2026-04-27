import { useState } from "react"

function App() {
  const [count, setCount] = useState(0)
    function Increment(){
  setCount(count+1)
}
  return (
    <>
    <button onClick={Increment}>{count}</button>
    <br/>
    {Array.from({ length: count }, (_, i) => (
      <h1 key={i}>hello</h1>
    ))}
    </>
  )
}
export default App
