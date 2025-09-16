interface CounterCardProps {
    label: string;
    value: number;
    min: number;
    max: number;
    onChange: (newValue: number) => void;
}
