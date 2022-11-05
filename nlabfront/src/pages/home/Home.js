import "./home.scss";
import React, {useEffect, useState} from 'react';
import gameLogo from './game_logo.png'
import Button from 'react-bootstrap/Button';
import {useSelector} from "react-redux";
import {jwtUtils} from "../../utils/jwtUtils";
import {Link} from "react-router-dom";
import api from "../../utils/api";
import {dailyCheck} from "../../components/DailyChecker";

const Home = () => {

    const [check, setCheck] = useState({message: ''});
    const [count, setCount] = useState({message: ''})
    const [reward, setReward] = useState({data:''})
    // function dailyCheckClick(){
    //     dailyCheck().then(response => {
    //         console.log(response)
    //         setstate
    //     })
    // }

    const dailyCheckClick = async () => {
        await api.get('/api/admin/daily_check', {
            headers: {
                Authorization: `Bearer ${token}`
            }
        }).then(response => {
            setCheck(response.data);
            console.log(response.data)
        })
    }

    const countingCheck = async () => {
        await api.get('/api/admin/counting_check', {
            headers: {
                Authorization: `Bearer ${token}`
            }
        }).then(response => {
            setCount(response.data);
            console.log(response.data)
        })
    }

    const compensation = async () => {
        await api.get('/api/admin/compensation', {
            headers: {
                Authorization: `Bearer ${token}`
            }
        }).then(response => {
            setReward(response.data);
            console.log(response.data)
        })
    }

    const token = useSelector(state => state.Auth.token);
    const [isAuth, setIsAuth] = useState(false);
    useEffect(() => {
        if (jwtUtils.isAuth(token)) {
            setIsAuth(true);
        } else {
            setIsAuth(false);
        }
    }, [token]);
    return (
        <div>

            <div className="home-wrapper">
                <div className="home-title">
                    <img className={"logo"} src={gameLogo}/>
                </div>
                <div className="home-contents">
                    {isAuth ? (
                        <>
                            <Button className={"loginButton"} variant={"warning"} onClick={dailyCheckClick}>출석체크</Button>
                            <p className={'checkMessage'}>
                                {check.message}
                            </p>
                            <Button className={"rewardButton"} variant={"info"} onClick={countingCheck}>출석확인</Button>
                            <p>
                                {count.message}
                            </p>
                            <Button className={"joinButton"} variant={"info"} onClick={compensation}>보상확인</Button>
                            <p>
                                {reward.data}
                            </p>
                        </>
                    ) : (
                        <>
                            <Link to="/login"><Button className={"loginButton"} variant={"warning"}>LOGIN</Button></Link>
                            <Link to={"/sign-up"}><Button className={"joinButton"} variant={"info"}>JOIN</Button></Link>
                        </>
                    )}
                    {/*<Button className={"loginButton"} variant={"warning"}>LOGIN</Button>*/}
                    {/*<Button className={"joinButton"} variant={"info"}>JOIN</Button>*/}
                </div>
                <div className="my-website">
                    <div className="my-website-title">Website</div>
                    <a href="https://www.youtube.com/channel/UC8h3CeUCar8Z4ux16bIMpCg" target="_blank">
                        YOUTUBE
                    </a>
                    <a href="https://ko.wikipedia.org/wiki/%EB%8C%80%ED%95%9C%EC%98%88%EC%88%98%EA%B5%90%EB%B3%B5%EC%9D%8C%EA%B5%90%ED%9A%8C"
                       target="_blank">
                        WIKI
                    </a>
                </div>
            </div>
        </div>
    )
}
export default Home;

