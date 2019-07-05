import React, { Component } from 'React';
import { View, Text} from 'react-native';
import {
    createBottomTabNavigator,
    createStackNavigator,
    createAppContainer
} from "react-navigation";
import Icon from 'react-native-vector-icons/AntDesign';

import HomeScreen from '../Views/Home'
import ReleaseScreen from '../Views/Release'
import ContactScreen from '../Views/Contact'
import UserScreen from '../Views/User'

class Test extends Component {
    render() {
        return (
            <View >
                <Text>
                    test page!
                </Text>
            </View>
        )
    }
}

const HomeStack = createStackNavigator({
    Home: { screen: HomeScreen },
    Test: { screen: Test},
});

const ReleaseStack = createStackNavigator({
    Release: { screen: ReleaseScreen },
    Test: { screen: Test},
});

const ContactStack = createStackNavigator({
    Contact: { screen: ContactScreen },
    Test: { screen: Test},
});

const UserStack = createStackNavigator({
    User: { screen: UserScreen },
    Test: { screen: Test},
});

const TabBar = createBottomTabNavigator({
    Home: { screen: HomeStack },
    Release: { screen: ReleaseStack },
    Contact: { screen: ContactStack },
    User: { screen: UserStack },
},{
    defaultNavigationOptions: ({ navigation }) => ({
        tabBarIcon: ({ focused, tintColor }) => {
            const { routeName } = navigation.state;
            let iconName;
            if (routeName === 'Home') {
                iconName = 'home'
            }
            else if (routeName === 'Release') {
                iconName = 'form'
            }
            else if (routeName === 'Contact') {
                iconName = 'message1'
            }
            else if (routeName === 'User') {
                iconName = 'user'
            }
            return <Icon name={iconName} size={25} color={tintColor} />;
        }
    })
});

const AppStack = createStackNavigator({
    Tabs: TabBar,
    TestPage: Test,
}, {
    defaultNavigationOptions: {
        headerStyle: {
            backgroundColor: '#ff9800',
        },
        headerTintColor: '#fff',
    }
});

export default createAppContainer(TabBar);

export { Test };