interface CalculateButtonProps {
    buttonName: string;
}

const CalculateButton = (props: CalculateButtonProps) => {
    return (
        <>
            <button
                className="bg-[#E83D67] w-full h-full text-white font-bold text-3xl "
                type="submit"
                formMethod="POST"
            >
                {props.buttonName}
            </button>
        </>
    );
};

export default CalculateButton;
