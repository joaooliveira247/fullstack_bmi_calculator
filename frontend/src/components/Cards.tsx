interface CardProps {
    height: number;
    children: React.ReactNode;
}

const Card = ({ height, children }: CardProps) => {
    return (
        <div
            className="bg-[#2D3249] w-full rounded-2xl flex justify-center items-center"
            style={{ height }}
        >
            {children}
        </div>
    );
};

export default Card;
