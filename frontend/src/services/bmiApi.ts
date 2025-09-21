export interface BMIData {
    bmi: number;
    status: string;
    message: string;
}

export async function getBMI(
    weight: number,
    height: number
): Promise<BMIData | undefined> {
    try {
        const response = await fetch("http://localhost:8000/bmi/", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ weight: weight, height: height }),
        });
        const data = await response.json();
        return data;
    } catch (e) {
        console.log(e);
        throw e;
    }
}
