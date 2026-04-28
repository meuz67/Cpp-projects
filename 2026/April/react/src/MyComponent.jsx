import { useState } from "react";
function MyComponent() {
    const [someValue, setSomeValue] = useState(0);

    return (
        <div>
            <input type = "checkbox" onChange={() => setSomeValue(someValue+1)}></input>
            {Array.from({ length: someValue }, (_, i) => (
                <p key={i}>{someValue}</p>
            ))}
        </div>
    );
}
export default MyComponent;
