import { BMIData } from "@/services/bmiApi";
import BaseLayout from "./BaseLayout";
import Card from "@/components/Cards";

interface ResultCardProps {
    bmi: BMIData;
    onClick: () => void;
}

const ResultCard = (props: ResultCardProps) => {
    return (
        <BaseLayout buttonName="Re - Calculate" onClickButton={props.onClick}>
            <h2 className="text-white font-semibold text-4xl">Your Result</h2>
            <Card height={521}>
                <h3 className="font-semibold text-4xl">{props.bmi.status}</h3>
                <h2 className="text-white font-semibold text-8xl">
                    {props.bmi.bmi}
                </h2>
                <p className="text-[#8B8C9E] font-semibold text-[16px] text-center">
                    {props.bmi.message}
                </p>
            </Card>
        </BaseLayout>
    );
};

export default ResultCard;
