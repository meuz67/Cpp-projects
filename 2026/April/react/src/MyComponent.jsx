import {useState} from "react";
function MyComponent() {
    const [name, setName] = useState("");
    const[quantity, setQuantity] = useState(0);
    return (
         <div>
            <input value = {name} onChange={(e) => setName(e.target.value)} placeholder="Enter name" />
            
            <p>Name: {name}</p>
         </div>
    )
}
export default MyComponent;