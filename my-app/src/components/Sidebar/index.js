import React from 'react'
import { SidebarContainer, Icon, CloseIcon, SidebarWrapper, SidebarMenu, SidebarLink, SideBtnWrap, SidebarRoute } from './elements'

const Sidebar = () => {
    return (
        <SidebarContainer>
            <Icon>
                <CloseIcon />
            </Icon>
            <SidebarWrapper>
                <SidebarMenu>
                    <SidebarLink to="home">Home</SidebarLink>
                    <SidebarLink to="resources">Resources</SidebarLink>
                    <SidebarLink to="about">About</SidebarLink>
                    <SidebarLink to="profile">Profile</SidebarLink>
                </SidebarMenu>
                <SideBtnWrap>
                    <SidebarRoute>Sign Up</SidebarRoute>
                </SideBtnWrap>
            </SidebarWrapper>
        </SidebarContainer>
    )
}

export default Sidebar;
