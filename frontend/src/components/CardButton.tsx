interface CardButtonProps {
  height: number;
  value: string;
  name: string;
  children: React.ReactNode;
}

const CardButton = ({ height, value, name, children }: CardButtonProps) => {
  return (
    <label
      className="cursor-pointer flex-1"
      style={{ height }}
    >
      <input
        type="radio"
        name={name}
        value={value}
        className="peer hidden"
      />
      <div
        className="w-full h-full rounded-2xl flex flex-col items-center justify-center gap-2 
                   bg-[#2D3249] peer-checked:bg-[#24263B] hover:bg-[#3A415C] transition"
      >
        {children}
      </div>
    </label>
  );
};

export default CardButton;
