import { useState } from "react";
import "./MyComponent.css"
function MyComponent() {
    function handleColorChange(e){
        setColor(e.target.value)
    }
    const [color, setColor] = useState("#FFFFF")
    return (
        <div>
            <div className="cube" style={{ backgroundColor: color }}></div>
            <div className="colorInput">
                <input type="color" onChange={handleColorChange} value={color} />
            </div>
        </div>
    );
}
export default MyComponent;
