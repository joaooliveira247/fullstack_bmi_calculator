import Card from "./Cards";

interface HeightCardProps {
    value: number;
    min?: number;
    max?: number;
    onChange: (newValue: number) => void;
}

const HeightCard = ({
    value,
    min = 50,
    max = 260,
    onChange,
}: HeightCardProps) => {
    return (
        <Card height={189}>
            <h2 className="text-2xl text-[#8B8C9E]">Height</h2>
            <div className="text-white">
                <span className="text-5xl">{value}</span>
                <span>cm</span>
            </div>
            <input
                type="range"
                min={min}
                max={max}
                className="h-[2px] w-[270px] accent-[#E83D67]"
                onChange={(e) => onChange(Number(e.target.value))}
            />
        </Card>
    );
};

export default HeightCard;
