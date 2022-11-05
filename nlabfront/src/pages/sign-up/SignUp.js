import axios from "axios";
import {toast, ToastContainer} from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import {Formik, ErrorMessage} from "formik";
import * as Yup from "yup";
import {Button, TextField} from "@mui/material";
import {useNavigate} from "react-router-dom";
import "./signUp.scss";

const SignUp = () => {
    const navigate = useNavigate();
    const validationSchema = Yup.object().shape({
        email: Yup.string()
            .email("올바른 이메일 형식이 아닙니다")
            .required("이메일을 입력하세요."),
        // username: Yup.string()
        //     // .min(2, "닉네임은 최소 2글자 이상입니다.")
        //     // .max(10, "닉네임은 최대 10글자입니다.")
        //     // .matches(
        //     //     /^[가-힣a-zA-Z][^!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?\s]*$/,
        //     //     "닉네임에 특수문자가 포함되면 안되고 숫자로 시작하면 안됩니다."
        //     // )
        //     .required("닉네임은 랜덤으로 생성됩니다."),
        password: Yup.string()
            .min(8, "비밀번호는 최소 8자리 이상입니다.")
            .max(16, "비밀번호는 최대 16자리입니다.")
            .required("패스워드를 입력하세요.")
            .matches(
                /^(?=.*[a-zA-Z])(?=.*[0-9])(?=.*[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?])[^\s]*$/,
                "알파벳, 숫자, 공백을 제외한 특수문자를 모두 포함해야 합니다."
            ),
        password2: Yup.string()
            .oneOf([Yup.ref("password"), null], "비밀번호가 일치하지 않습니다.")
            .required("필수 입력 값입니다."),
    });
    const submit = async (values) => {
        const {email, password} = values;
        try {
            fetch('/api/register', {
                method: "POST",
                headers: {
                    "content-Type": "application/json",
                },
                body: JSON.stringify({
                    email,
                    password,
                })
            })
            toast.success(<h3>회원가입이 완료되었습니다.<br/>로그인 하세요.</h3>, {
                position: "top-center",
                autoClose: 2000
            });

        } catch (e) {
            // 서버에서 받은 에러 메시지 출력
            toast.error(e.response.data.message + "Error", {
                position: "top-center",
            });
        }
        setTimeout(() => {
            navigate("/login");
        }, 2000);
    };

    return (
        <Formik
            initialValues={{
                email: "",
                password: "",
                password2: "",
            }}
            validationSchema={validationSchema}
            onSubmit={submit}
            validateOnMount={true}
        >
            {({values, handleSubmit, handleChange, errors}) => (
                <div className="signup-wrapper">
                    <ToastContainer/>
                    <form onSubmit={handleSubmit} autoComplete="off">
                        <div className="input-forms">
                            <div className="input-forms-item">
                                <div className="input-label">이메일</div>
                                <TextField
                                    value={values.email}
                                    name="email"
                                    className={"email"}
                                    variant="outlined"
                                    onChange={handleChange}
                                />
                                <div className="error-message">
                                    {errors.email}
                                </div>
                            </div>
                            <div className="input-forms-item">
                                <div className="input-label">닉네임</div>
                                <div className="error-message">
                                    닉네임은 자동으로 생성됩니다.
                                </div>
                            </div>
                            <div className="input-forms-item">
                                <div className="input-label">비밀번호</div>
                                <TextField
                                    value={values.password}
                                    name="password"
                                    className={"password"}
                                    variant="outlined"
                                    type="password"
                                    onChange={handleChange}
                                />
                                <div className="error-message">
                                    {errors.password}
                                </div>
                            </div>
                            <div className="input-forms-item">
                                <div className="input-label">비밀번호 확인</div>
                                <TextField
                                    value={values.password2}
                                    name="password2"
                                    className={"password"}
                                    variant="outlined"
                                    type="password"
                                    onChange={handleChange}
                                />
                                <div className="error-message">
                                    {errors.password2}
                                </div>
                            </div>
                            <Button
                                color="primary"
                                variant="contained"
                                fullWidth
                                type="submit"
                            >
                                회원가입
                            </Button>
                        </div>
                    </form>
                </div>
            )}
        </Formik>
    );
};

export default SignUp;
