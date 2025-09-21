"use client";
import { getBMI } from "@/services/bmiApi";
import Home from "@/widgets/Home";
import { useState } from "react";

export default function Page() {
    const [height, setHeight] = useState(152);
    const [weight, setWeight] = useState(60);
    const [age, setAge] = useState(26);

    const handleCalculate = async () => {
        try {
            const result = await getBMI(weight, height);
            console.log(result);
        } catch (err) {
            console.error(err);
        }
    };

    return (
        <Home
            height={height}
            setHeight={setHeight}
            weight={weight}
            setWeight={setWeight}
            age={age}
            setAge={setAge}
            onClickButton={handleCalculate}
        />
    );
}
