import React, { Component } from 'react';
import {View, Text, StyleSheet, ScrollView, TouchableOpacity, Image, Dimensions, Alert} from 'react-native';
import HTTP from '../../Network/Network';
import {parseStatus, parseTimeStamp} from "../../Utils/ParseInfo";
import MyAvatar from '../../Components/MyAvatar';
import Video from "react-native-video";
import Config from "../../Config";
import {Avatar, Button} from "react-native-elements";
import {NavigationActions} from "react-navigation";

const {width, height, scale} = Dimensions.get('window');

class GoodInfoImage extends Component {
    render () {
        let {fileID} = this.props;
        let imageUri = (Config.fetchPrefix + 'file/' + fileID);
        let imageWidth=width * 0.85, imageHeight=width * 0.85;
        /*Image.getSize(imageUri)
            .then((width, height) => {
                imageWidth = 0.85 * width;
                imageHeight = height * imageWidth / width;
            })
            .catch((error) => {
                console.log(error); //取宽，高出错
            });*/
        return (
            <TouchableOpacity activeOpacity={0.7} onLongPress={() => {}}>
                <Image style={{
                    width: imageWidth,
                    height: imageHeight,
                    resizeMode: 'contain',
                    borderColor: 'white',
                    borderWidth: 1,
                }} source={{uri: imageUri}} />
            </TouchableOpacity>
        );
    }
}

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
    constructor(props) {
        super(props);
        this.state = {
            loaded: false,
            contentToken: '',
            files: [],
        };
        this.keyID = 0;
    };

    componentDidMount() {
        this.fetchData();
    }

    fetchData() {
        //console.warn(Config.userInfo);
        //let obj = { userID: Config.userInfo.userID };
        let obj = {  };
        let {contentID} = this.props;
        HTTP.get(('/content/' + contentID), obj)
            .then((response) => {
                //console.warn(response);
                this.setState({
                    contentToken: response.contentToken,
                    files: response.files,
                    loaded: true,
                });
            })
            .catch((error) => console.error(error));
    }

    renderVideo(video) {
        console.log('rendering video');
        return (<View style={{height: 300, width: 300}}>
            <Video source={{uri: video.uri, type: video.mime}}
                   style={{position: 'absolute',
                       top: 0,
                       left: 0,
                       bottom: 0,
                       right: 0
                   }}
                   rate={1}
                   paused={false}
                   volume={1}
                   muted={false}
                   resizeMode={'cover'}
                   onError={e => console.log(e)}
                   onLoad={load => console.log(load)}
                   repeat={true} />
        </View>);
    }

    renderImage = (fileID) => {
        //console.warn(fileID);
        return (
            <GoodInfoImage fileID={fileID}/>
        )
    };

    renderAsset(file) {
        //console.warn(file);
        if (file.type === 2) {
            return this.renderVideo(file.fileID);
        }
        return this.renderImage(file.fileID);
    }

    render() {
        if (this.state.loaded === false) {
            return (
                <View style={styles.container}>
                    <Text>图片加载中...</Text>
                </View>
            );
        } else {
            console.warn(this.state.files);
            return (
                <View style={{alignItems: 'center'}}>
                    <View style={{backgroundColor: '#FCFCFC'}}>
                        {this.state.files.length > 0 ? this.state.files.map(file => <View key={this.keyID++}>{this.renderAsset(file)}</View>) : null}
                    </View>
                </View>
            );
        }
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
                <View style={{backgroundColor: '#FCFCFC'}}>
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
            let ReserveButtonColor = this.state.status === 1 ? 'red' : 'grey';
            return (
                <ScrollView>
                    <View style={{backgroundColor: '#EFEFF5'}}>
                        <InfoPart infoType={this.infoType} item={this.state}/>
                        <View style={{height: 20, backgroundColor: '#EFEFF5'}}/>
                        {this.state.contentID === undefined? null : <ImagePart contentID={this.state.contentID}/>}
                        <View style={{height: 20, backgroundColor: '#EFEFF5'}}/>
                        <UserPart userID={this.state.userID} navigate={(params) => this.navigate(params)}/>
                        <View style={{height: 20, backgroundColor: '#EFEFF5'}}/>
                        <Button
                            title='申请预约'
                            titleStyle={{color: 'white', fontSize: 17}}
                            buttonStyle={{backgroundColor: ReserveButtonColor}}
                            containerStyle={{width: 160, marginLeft: (width / 2 - 80)}}
                            raised={true}
                            onPress={() => {
                                if (this.state.status === 1) {
                                    Alert.alert(
                                        '成功预约！',
                                        '您已经成功预约了该交易信息，请等待对方答复',
                                        [
                                            {text: '好', onPress: () => {
                                                    this.setState({isChangeTelephoneVisible: true})
                                                }}
                                        ],
                                        {cancelable: false},
                                    )
                                }
                            }}
                        />
                    </View>
                    <View style={{height: 20, backgroundColor: '#EFEFF5'}}/>
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
