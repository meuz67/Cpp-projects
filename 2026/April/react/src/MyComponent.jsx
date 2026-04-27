import { useState } from "react";
import "./App.css"
function MyComponent() {
    const [count, setCount] = useState(0);
    return(
        <div class="stuff">
            <h1>Count: {count}</h1>
            <button onClick={() => setCount(count + 1)}>Increment</button>
            <button onClick={()=>setCount(count - 1 )}>Decrement</button>
            <button onClick={()=>setCount(0)}>Reset</button>
        </div>
    )
}
export default MyComponent;