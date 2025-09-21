import { Dispatch, SetStateAction } from "react";
import BaseLayout from "./BaseLayout";
import CardButton from "@/components/CardButton";
import { IconGenderFemale, IconGenderMale } from "@/assets/icons";
import HeightCard from "@/components/HeightCard";
import CounterCard from "@/components/CounterCard";

type StateSetter = Dispatch<SetStateAction<number>>;

interface HomeProps {
    height: number;
    setHeight: StateSetter;
    weight: number;
    setWeight: StateSetter;
    age: number;
    setAge: StateSetter;
    onClickButton: () => void;
}

export default function Home({
    height,
    setHeight,
    weight,
    setWeight,
    age,
    setAge,
    onClickButton,
}: HomeProps) {
    return (
        <BaseLayout buttonName="Calculate" onClickButton={onClickButton}>
            <div className="flex gap-[9px]">
                <CardButton height={180} name="gender" value="male">
                    <IconGenderMale
                        aria-label="male gender icon"
                        size={144}
                        color="white"
                    />
                    <h2 className="text-2xl text-[#8B8C9E]">Male</h2>
                </CardButton>
                <CardButton height={180} name="gender" value="female">
                    <IconGenderFemale
                        aria-label="male gender icon"
                        size={144}
                        color="white"
                    />
                    <h2 className="text-2xl text-[#8B8C9E]">Female</h2>
                </CardButton>
            </div>

            <div className="w-full">
                <HeightCard value={height} onChange={setHeight}></HeightCard>
            </div>

            <div className="flex gap-[9px]">
                <CounterCard
                    label="Weight"
                    min={4}
                    max={635}
                    value={weight}
                    onChange={setWeight}
                />
                <CounterCard
                    label="Age"
                    min={1}
                    max={120}
                    value={age}
                    onChange={setAge}
                />
            </div>
        </BaseLayout>
    );
}
