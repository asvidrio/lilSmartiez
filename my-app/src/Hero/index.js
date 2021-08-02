import React, {useState} from 'react'
import Video from '../videos/video.mp4'
import { HeroContainer, HeroBg, VideoBg, HeroContent, HeroH1, HeroP, HeroBtnWrapper, ArrowForward, ArrowRight } from './elements'
import {Button} from '../ButtonElement'

const Hero = () => {
    const [hover, setHover] = useState(false)

    const onHover = () => {
        setHover(!hover)
    }

    return (
        <HeroContainer id="Home">
            <HeroBg>
                <VideoBg autoPlay loop muted src={Video} type='video/mp4' />
            </HeroBg>
            <HeroContent>
                <HeroH1>
                    Providing Resources for Parents/Teachers to help their children gain access to STEM-education
                </HeroH1>
                <HeroP>
                    Create an account to keep track of your saved resources.
                </HeroP>
                <HeroBtnWrapper>
                    <Button to='signin' onMouseEnter={onHover} onMouseLeave={onHover} primary='true' dark='false'>
                        Get started {hover ? <ArrowForward /> : <ArrowRight/>}
                    </Button>
                </HeroBtnWrapper>
            </HeroContent>
        </HeroContainer>
    )
}

export default Hero
