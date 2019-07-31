import React, { Component } from 'react';
import {View, Text, StyleSheet, ScrollView, TouchableOpacity} from 'react-native';
import HTTP from '../../Network/Network';
import {parseStatus, parseTimeStamp} from "../../Utils/ParseInfo";
import MyAvatar from '../../Components/MyAvatar';

class InfoPart extends Component {
    render() {
        let { infoType, item } = this.props;
        let NAME = infoType === 'sellInfo' ? '出售' : '求购';
        let STATUS = infoType === 'sellInfo' ? '出售价格：' : '预期价格：';
        return (
            <View style={{marginTop: 20}}>
                <View style={{backgroundColor: '#FCFCFC'}}>
                    <Text style={{marginLeft: 20, marginRight: 20, marginTop: 20}} selectable={true}>
                        <Text style={[styles.headerText, styles.goodType]}>{NAME}</Text>
                        <Text style={styles.headerText}>物品：</Text>
                        <Text style={styles.ratingText}>{item.goodName}</Text>
                    </Text>
                </View>
                <View style={{backgroundColor: '#FCFCFC'}}>
                    <Text style={{marginLeft: 20, marginRight: 20, marginTop: 10}}>
                        <Text style={styles.headerText}>具体描述：</Text>
                        <Text style={styles.ratingText}>{item.description}</Text>
                    </Text>
                </View>
                <View style={{backgroundColor: '#FCFCFC'}}>
                    <Text style={{marginLeft: 20, marginRight: 20, marginTop: 10}}>
                        <Text style={styles.headerText}>信息状态：</Text>
                        <Text style={styles.ratingText}>{parseStatus(item.status)}</Text>
                    </Text>
                </View>
                <View style={{backgroundColor: '#FCFCFC'}}>
                    <Text style={{marginLeft: 20, marginRight: 20, marginTop: 10}}>
                        <Text style={styles.headerText}>{STATUS}</Text>
                        <Text style={styles.ratingText}>￥{item.price}</Text>
                    </Text>
                </View>
                <View style={{backgroundColor: '#FCFCFC'}}>
                    <Text style={{marginLeft: 20, marginRight: 20, marginTop: 10}}>
                        <Text style={styles.headerText}>发布时间：</Text>
                        <Text style={styles.ratingText}>{parseTimeStamp(item.releaseTime)}</Text>
                    </Text>
                </View>
                <View style={{backgroundColor: '#FCFCFC'}}>
                    <Text style={{marginLeft: 20, marginRight: 20, marginTop: 10, marginBottom: 20}}>
                        <Text style={styles.headerText}>物品标签：</Text>
                        <Text style={styles.ratingText}>暂无</Text>
                    </Text>
                </View>
            </View>
        );
    }
}

class ImagePart extends Component {
    render() {
        let {contentID} = this.props;
        return (
            <View style={{marginTop: 20}}>
                <View style={{backgroundColor: '#FCFCFC'}}>
                    <Text style={{marginLeft: 20, marginRight: 20}} selectable={true}>
                        <Text style={styles.headerText}>图片部分{contentID}</Text>
                    </Text>
                </View>
                <View style={{height: 10, backgroundColor: '#EFEFF5'}} />
            </View>
        )
    }
}

class UserPart extends  Component {
    constructor(props) {
        super(props);
        this.state = {
            loaded: false,
            userID: -1,
            userName: '',
            avatarID: '',
        };
    };

    componentDidMount() {
        this.fetchData();
    }

    fetchData() {
        //console.warn(Config.userInfo);
        //let obj = { userID: Config.userInfo.userID };
        let obj = {  };
        let {userID} = this.props;
        //console.warn(obj);
        //console.warn("http://202.120.40.8:30711/v1"+'/sellInfo/' + this.sellInfoID);
        HTTP.get(('/user/' + userID), obj)
            .then((response) => {
                //console.warn(response);
                this.setState({
                    userID: response.userID,
                    userName: response.userName,
                    avatarID: response.avatarID,
                    loaded: true,
                });
            })
            .catch((error) => console.error(error));
    }

    render() {
        if (this.state.loaded === false) {
            return (
                <View style={styles.container}>
                    <Text>用户信息加载中...</Text>
                </View>
            )
        }
        else {
            let {userID, navigate} = this.props;
            let userName = this.state.userName;
            //console.warn(userName);
            return (
                <View style={{backgroundColor: '#FCFCFC', marginBottom: 20}}>
                    <View style={{
                        backgroundColor: '#FCFCFC',
                        flexDirection: 'row',
                        justifyContent: 'space-between',
                        marginTop: 20,
                        marginBottom: 20,
                        marginLeft: 20,
                        marginRight: 20,
                    }}>
                        <View style={{marginLeft: 10}}>
                            <Text style={styles.headerText}>信息发布人：</Text>
                            <Text style={{color: 'black', fontSize: 20}}> {this.state.userName}</Text>
                            <TouchableOpacity onPress={() => {navigate({ userID, userName });}}>
                                <Text style={styles.headerText}>点击此处查看TA的更多信息</Text>
                            </TouchableOpacity>
                        </View>
                        <MyAvatar avatarID={this.state.avatarID}/>
                    </View>
                    <View style={{height: 10, backgroundColor: '#EFEFF5'}}/>
                </View>
            )
        }
    }
}

export default class GoodInfoScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            loaded: false,
            status: -1,
            releaseTime: -1,
            validTime: -1,
            goodName: '',
            price: -1,
            description: '',
            contentID: '',
            userID: -1,
        };
        this.params = this.props.navigation.state.params;
        //console.warn(this.params)
        this.infoType = this.params? this.params.infoType : null;
        this.infoID = this.infoType === 'sellInfo' ? this.params.sellInfoID : this.params.buyInfoID;
        this.fetchData = this.fetchData.bind(this);
    };

    componentDidMount() {
        this.fetchData();
    }

    fetchData() {
        //console.warn(Config.userInfo);
        //let obj = { userID: Config.userInfo.userID };
        let obj = {  };
        //console.warn(obj);
        //console.warn("http://202.120.40.8:30711/v1"+'/sellInfo/' + this.sellInfoID);
        HTTP.get(('/' + this.infoType + '/' + this.infoID), obj)
            .then((response) => {
                //console.warn(response);
                this.setState({
                    status: response.status,
                    releaseTime: response.releaseTime,
                    validTime: response.validTime,
                    goodName: response.goodName,
                    price: response.price,
                    description: response.description,
                    contentID: response.contentID,
                    userID: response.userID,
                    loaded: true,
                });
            })
            .catch((error) => console.error(error));
    }

    static navigationOptions = ({navigation}) => ({
        headerTitle: (<Text numberOfLines={1} style={{flex:1, color: '#298BFF', fontSize: 23}}>{navigation.state.params.header}</Text>)
    });

    navigate = (params) => {
        //console.warn(params);
        this.props.navigation.push('UserInfoForOthers', params);
    };

    render() {
        if (this.state.loaded === false) {
            return (
                <View style={styles.container}>
                    <Text>加载中...</Text>
                </View>
            )
        }
        else {
            //console.warn(this.state.SellInfoList[0].releaseTime);
            //console.warn(TimeStamptoDate(this.state.SellInfoList[0].releaseTime));
            //console.warn(TimeStampNow());
            //console.warn(TimeStamptoDate(TimeStampNow()));
            return (
                <ScrollView>
                    <View style={{backgroundColor: '#EFEFF5'}}>
                        <InfoPart infoType={this.infoType} item={this.state}/>
                        <ImagePart contentID={this.state.contentID}/>
                        <UserPart userID={this.state.userID} navigate={(params) => this.navigate(params)}/>
                    </View>
                </ScrollView>
            )
        }
    }
}

let styles = StyleSheet.create({
    container: {
        flex: 1,
        flexDirection: "row",
        justifyContent: "center",
        alignItems: "center",
        backgroundColor: "#EFEFF5"
    },
    subtitleView: {
        //flexDirection: 'row',
        paddingLeft: 10,
        paddingTop: 5
    },
    ratingText: {
        color: '#042A2B',
        fontSize: 18,
    },
    headerText: {
        color: '#5EB1BF',
        fontSize: 18,
    },
    goodType: {
        fontSize: 23,
        fontWeight: 'bold',
    }
});
