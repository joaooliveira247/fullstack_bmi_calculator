import CalculateButton from "@/components/Button";

interface BaseLayoutProps {
    children: React.ReactNode;
}

const BaseLayout = (props: BaseLayoutProps) => {
    return (
        <div className="h-screen flex justify-center items-center bg-white md:bg-gray-100">
            <div className="w-full h-full md:w-[360px] md:h-[800px] md:rounded-2xl overflow-hidden flex flex-col">
                {/* Header */}
                <div className="text-center bg-[#24263B]">
                    <h1 className="font-bold p-[13px] text-white">
                        BMI Calculator
                    </h1>
                </div>
                {/* Content */}
                <div className="flex-1 p-[20px] bg-[#1C2135] flex flex-col gap-3">
                    {props.children}
                </div>
                <div className="flex-1">
                    <CalculateButton buttonName="Calculate" />
                </div>
            </div>
        </div>
    );
};

export default BaseLayout;
