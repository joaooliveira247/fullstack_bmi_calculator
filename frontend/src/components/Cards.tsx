interface CardProps {
    height: number;
    children: React.ReactNode;
}

const Card = (props: CardProps) => {
    const cardProps = `bg-[#2D3249] w-full h-[${props.height}px] rounded-2xl`;
    return <div className={cardProps}>{props.children}</div>;
};

export default Card;
