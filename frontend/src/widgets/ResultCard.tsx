import { BMIData } from "@/services/bmiApi";
import BaseLayout from "./BaseLayout";
import Card from "@/components/Cards";

type BMIStatus =
    | "Severe thinness"
    | "Moderate thinness"
    | "Mild thinness"
    | "Normal"
    | "Overweight"
    | "Obese class I"
    | "Obese class II"
    | "Obese class III";

const bmiStatusTextColors: Record<BMIStatus, string> = {
    "Severe thinness": "text-red-800",
    "Moderate thinness": "text-orange-700",
    "Mild thinness": "text-yellow-700",
    Normal: "text-green-600",
    Overweight: "text-amber-700",
    "Obese class I": "text-rose-700",
    "Obese class II": "text-pink-700",
    "Obese class III": "text-purple-700",
};

interface ResultCardProps {
    bmi: BMIData | null;
    onClick: () => void;
}

const ResultCard = (props: ResultCardProps) => {
    if (props.bmi) {
        return (
            <BaseLayout
                buttonName="Re - Calculate"
                onClickButton={props.onClick}
            >
                <h2 className="text-white font-semibold text-4xl">
                    Your Result
                </h2>
                <Card height={521}>
                    <h3
                        className={`font-semibold text-4xl ${
                            bmiStatusTextColors[props.bmi.status as BMIStatus]
                        }`}
                    >
                        {props.bmi.status}
                    </h3>
                    <h2 className="text-white font-semibold text-8xl">
                        {props.bmi.BMI}
                    </h2>
                    <p className="text-[#8B8C9E] font-semibold text-[16px] text-center">
                        {props.bmi.message}
                    </p>
                </Card>
            </BaseLayout>
        );
    }
};

export default ResultCard;
