import { Dispatch, SetStateAction } from "react";

type StateSetter = Dispatch<SetStateAction<number>>;

interface HomeProps {
    height: number;
    setHeight: StateSetter;
    weight: number;
    setWeight: StateSetter;
    age: number;
    setAge: StateSetter;
    onClickButton: () => void;
}
