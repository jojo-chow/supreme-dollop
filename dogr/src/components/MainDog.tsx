import { useState } from "react"

/**
 * Renders the dog image that was uploaded or nothing on initial state
 * @returns image or placeholder
 */
const MainDog = () => {
    const [mainDog, setMainDog] = useState(false);

    const toggleMainDog = (event: React.MouseEvent<HTMLButtonElement>) => {
        setMainDog(!mainDog);
    }

    return (
        <div className="flex">
            Main Dog: {mainDog}
            <button onClick={toggleMainDog}>Toggle</button>
        </div>
    );
}

export default MainDog;