"use client";
import { BMIData, getBMI } from "@/services/bmiApi";
import Home from "@/widgets/Home";
import ResultCard from "@/widgets/ResultCard";
import { useState } from "react";

export default function Page() {
    const [height, setHeight] = useState(152);
    const [weight, setWeight] = useState(60);
    const [age, setAge] = useState(26);
    const [step, setStep] = useState<"form" | "result">("form");
    const [resultData, setResultData] = useState<null | BMIData>(null);

    const handleCalculate = async () => {
        const heightMeters = parseFloat((height / 100).toFixed(2));
        try {
            const result = await getBMI(weight, heightMeters);
            if (result) {
                setResultData(result);
                setStep("result");
            }
        } catch (err) {
            console.error(err);
        }
    };

    const handleBack = () => {
        setStep("form");
        setResultData(null);
    };

    return (
        <>
            {step === "form" && (
                <Home
                    height={height}
                    setHeight={setHeight}
                    weight={weight}
                    setWeight={setWeight}
                    age={age}
                    setAge={setAge}
                    onClickButton={handleCalculate}
                />
            )}
            {step === "result" && (
                <ResultCard bmi={resultData} onClick={handleBack} />
            )}
        </>
    );
}
