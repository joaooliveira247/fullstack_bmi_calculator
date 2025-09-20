interface CalculateButtonProps {
    buttonName: string;
    onClick: () => void;
}

const CalculateButton = (props: CalculateButtonProps) => {
    return (
        <>
            <button
                className="bg-[#E83D67] w-full h-full text-white font-bold text-3xl "
                type="submit"
                formMethod="POST"
                onClick={props.onClick}
            >
                {props.buttonName}
            </button>
        </>
    );
};

export default CalculateButton;
