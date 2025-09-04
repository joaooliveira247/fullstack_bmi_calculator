interface CalculateButtonProps {
    buttonName: string;
}

const CalculateButton = (props: CalculateButtonProps) => {
    return (
        <>
            <button
                type="submit"
                formMethod="POST"
                className="bg-[#E83D67] w-full h-[100] rounded-b-2xl"
            >
                <h2 className="text-3xl font-bold text-white">
                    {props.buttonName}
                </h2>
            </button>
        </>
    );
};

export default CalculateButton;
