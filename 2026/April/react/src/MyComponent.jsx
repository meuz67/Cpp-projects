import { useState } from "react";
function MyComponent() {
    const[firstValue, setFirstValue] = useState(0);
    const [secondValue, setSecondValue] = useState(0);
    return (
        <div>
            <input type="number" onChange={(e) => setFirstValue(e.target.value)} value={firstValue} />
            <input type="number" onChange={(e) => setSecondValue(e.target.value)} value={secondValue} />
            <p>{firstValue * secondValue}</p>
        </div>
    );
}
export default MyComponent;
