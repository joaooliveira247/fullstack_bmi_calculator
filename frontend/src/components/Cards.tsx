interface CardProps {
    height: number;
    children: React.ReactNode;
}

const Card = ({ height, children }: CardProps) => {
    return (
        <div
            className="bg-[#2D3249] w-full rounded-2xl flex flex-col items-center justify-center gap-2"
            style={{ height }}
        >
            {children}
        </div>
    );
};

export default Card;
