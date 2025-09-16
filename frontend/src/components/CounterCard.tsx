import { IconMinus, IconPlus } from "@/assets/icons";
import Card from "./Cards";

interface CounterCardProps {
    label: string;
    value: number;
    min: number;
    max: number;
    onChange: (newValue: number) => void;
}

const CounterCard = ({
    label,
    value,
    min,
    max,
    onChange,
}: CounterCardProps) => {
    const increase = () => value < max && onChange(value + 1);
    const decrease = () => value > min && onChange(value - 1);
    return (
        <Card height={180}>
            <h2 className="text-2xl text-white">{label}</h2>
            <span className="text-5xl text-white">{value}</span>
            <div className="flex gap-8">
                <button
                    type="button"
                    className="flex items-center justify-center w-10 h-10 rounded-full bg-[#8B8C9E] shadow-lg hover:bg-gray-600"
                    onClick={decrease}
                >
                    <IconMinus size={20} stroke={2} className="text-white" />
                </button>
                <button
                    type="button"
                    className="flex items-center justify-center w-10 h-10 rounded-full bg-[#8B8C9E] shadow-lg hover:bg-gray-600"
                    onClick={increase}
                >
                    <IconPlus size={20} stroke={2} className="text-white" />
                </button>
            </div>
        </Card>
    );
};

export default CounterCard;
