import React, { Component } from 'react';
import {
    Alert,
    Text,
    TextInput,
    View,
    StyleSheet,
    Dimensions,
    ScrollView,
    TouchableOpacity,
    Image,
    NativeModules
} from 'react-native';
import {Avatar, Badge, Button, ButtonGroup, SearchBar} from 'react-native-elements';
import Config from '../../Config';
import {NavigationActions} from "react-navigation";
import Textarea from 'react-native-textarea';
import {TimeStamptoDate, TimeStampNow, DatetoTimeStamp} from "../../Utils/TimeStamp";
import Video from "react-native-video";
import {isValidTimeValid, isPriceValid} from "../../Utils/CheckValidity";
import HTTP from "../../Network/Network";
import {isLogin} from "../../Config";

const {width, height, scale} = Dimensions.get('window');

let ImagePicker = NativeModules.ImageCropPicker;

class ReleaseImage extends Component {
    render () {
        //console.warn(this.props.image);
        return (
            <TouchableOpacity onLongPress={() => {this.props.deleteImage(this.props.index)}}>
                <Image style={{
                    width: (1.05 / 5.05 * width),
                    height: (1.05 / 5.05 * width),
                    resizeMode: 'contain',
                    borderColor: 'white',
                    borderWidth: 1,
                }} source={this.props.image} />
            </TouchableOpacity>
        )
    }
}

export default class ReleaseScreen extends Component {
    constructor(props) {
        super(props);
        this.state = {
            selectedIndex: 0,
            goodName: '',
            description: '',
            Price: '',
            validTime: '',
            images: [],
        };
        this.updateIndex = this.updateIndex.bind(this);
        this.keyID = 0;
        this.contentID = '';
        this.contentToken = '';
        this.uploadImageSuccess = false;
    };

    static navigationOptions = {
        headerTitle: (<Text style={{flex:1, color: '#298BFF', fontSize: 23, textAlign: 'center'}}>发布交易信息</Text>)
    };

    updateIndex (selectedIndex) {
        this.setState({
            selectedIndex: selectedIndex,
        })
    }

    updateGoodName = (goodName) => {
        this.setState({
            goodName: goodName,
        })
    };

    updateDescription = (description) => {
        this.setState({
            description: description,
        })
    };

    updatePrice = (Price) => {
        this.setState({
            Price: Price,
        })
    };

    updateValidTime = (validTime) => {
        this.setState({
            validTime: validTime,
        });
    };

    pickMultiple() {
        ImagePicker.openPicker({
            multiple: true,
            forceJpg: true,
        }).then(selectedImages => {
            this.keyID = 0;
            this.setState(previousState => {
                return {
                    images: previousState.images.concat(selectedImages.map(i => {
                        console.log('received image', i);
                        return {uri: i.path, width: i.width, height: i.height, mime: i.mime};
                    })),
                }});
        }).catch(e => {
            //console.warn('出错啦！');
            //console.warn(e);
        });
    }

    deleteImage = (index) => {
        //console.warn('删除图片' + index);
        Alert.alert(
            '即将删除图片',
            '确定删除该图片吗？',
            [
                {
                    text: '取消',
                    style: 'cancel',
                },
                {
                    text: '确定', onPress: () => {
                        this.keyID = 0;
                        this.setState({
                            images: this.state.images.filter((_, i) => i !== index),
                        });
                    }
                },
            ],
            {cancelable: false},
        );
    };

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

    renderImage = (image, index) => {
        return (
            <ReleaseImage image={image} index={index} deleteImage={index => this.deleteImage(index)}/>
        )
    };

    renderAsset(image, index) {
        if (image.mime && image.mime.toLowerCase().indexOf('video/') !== -1) {
            return this.renderVideo(image);
        }
        return this.renderImage(image, index);
    }

    render() {
        const buttons = ['出售信息', '求购信息'];
        const { selectedIndex } = this.state;
        if (isLogin() === false) {
            Alert.alert(
                '未登录',
                '无法在未登录状态下发布信息，是否切换至登录界面？',
                [
                    {
                        text: '取消',
                        onPress: () => {
                            this.props.navigation.navigate('Home');
                        },
                        style: 'cancel',
                    },
                    {
                        text: '确定', onPress: () => {
                            this.props.navigation.navigate('Login');
                        }
                    },
                ],
                {cancelable: false},
            );
            return (<View/>);
        } else {
            Config.isReleaseRender = true;
            return (
                <ScrollView>
                    <View style={{backgroundColor: '#EFEFF5'}}>
                        <ButtonGroup
                            onPress={this.updateIndex}
                            selectedIndex={selectedIndex}
                            buttons={buttons}
                            containerStyle={{height: 35}}
                            disabledTextStyle={{fontSize: 17}}
                            selectedTextStyle={{fontSize: 17, fontWeight: 'bold'}}
                        />
                        <View style={{flexDirection: 'row', backgroundColor: 'white'}}>
                            <Button
                                title='物品名称'
                                titleStyle={{color: 'steelblue', fontSize: 14, fontWeight: 'bold'}}
                                containerStyle={{flex: 1.1}}
                                buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                                raised={true}
                                onPress={() => {alert(this.state.goodName)}}
                            />
                            <SearchBar
                                searchIcon={false}
                                clearIcon={false}
                                cancelIcon={false}
                                lightTheme={true}
                                containerStyle={{
                                    flex: 4,
                                    height: 55,
                                    backgroundColor: 'white',
                                    borderWidth: 0, //no effect
                                    shadowColor: 'white', //no effect
                                    borderBottomColor: 'transparent',
                                    borderTopColor: 'transparent',
                                }}
                                inputContainerStyle={{
                                    marginTop: -9,
                                    height: 55,
                                    backgroundColor: 'white',
                                }}
                                placeholder={'必填。输入物品名称'}
                                onChangeText={ this.updateGoodName }
                                value={ this.state.goodName }
                            />
                        </View>
                        <View style={{height: 10}} />
                        <View style={{flexDirection: 'row', backgroundColor: 'white'}}>
                            <Button
                                title='详细描述'
                                titleStyle={{color: 'steelblue', fontSize: 14, fontWeight: 'bold'}}
                                containerStyle={{flex: 1.05}}
                                buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                                raised={true}
                                onPress={() => {alert(this.state.description)}}
                            />
                            <Textarea
                                containerStyle={{
                                    flex: 4,
                                    height: 150,
                                    backgroundColor: 'white',
                                    borderWidth: 0, //no effect
                                    shadowColor: 'white', //no effect
                                    borderBottomColor: 'transparent',
                                    borderTopColor: 'transparent',
                                }}
                                style={{marginLeft: 14, marginRight: 13, fontSize: 18, }}
                                onChangeText={this.updateDescription}
                                defaultValue={this.state.description}
                                maxLength={360}
                                placeholder={'选填。可输入对物品的详细描述或者列出物品清单'}
                                placeholderTextColor={'grey'}
                                underlineColorAndroid={'transparent'}
                            />
                        </View>
                        <View style={{height: 10}} />
                        <View>
                            <Text style={{alignItems: 'center', justifyContent: 'center'}}>
                                商品标签
                            </Text>
                        </View>
                        <View style={{height: 10}} />
                        <View style={{flexDirection: 'row'}}>
                            <TouchableOpacity onPress={this.pickMultiple.bind(this)}>
                                <Image style={{width: (1.05 / 5.05 * width), height: (1.05 / 5.05 * width), resizeMode: 'contain',}} source={require('../../assets/images/addPhotoVideo.jpg')}/>
                            </TouchableOpacity>
                            <ScrollView horizontal={true}>
                                {this.state.images.length > 0 ? this.state.images.map((image, index) => <View key={this.keyID++}>{this.renderAsset(image, index)}</View>) : null}
                            </ScrollView>
                        </View>
                        <View style={{height: 10}} />
                        <View style={{flexDirection: 'row', backgroundColor: 'white'}}>
                            <Button
                                title='预期价格'
                                titleStyle={{color: 'steelblue', fontSize: 14, fontWeight: 'bold'}}
                                containerStyle={{flex: 1.1}}
                                buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                                raised={true}
                                onPress={() => {alert(this.state.Price)}}
                            />
                            <SearchBar
                                searchIcon={false}
                                clearIcon={false}
                                cancelIcon={false}
                                lightTheme={true}
                                containerStyle={{
                                    flex: 4,
                                    height: 55,
                                    backgroundColor: 'white',
                                    borderWidth: 0, //no effect
                                    shadowColor: 'white', //no effect
                                    borderBottomColor: 'transparent',
                                    borderTopColor: 'transparent',
                                }}
                                inputContainerStyle={{
                                    marginTop: -9,
                                    height: 55,
                                    backgroundColor: 'white',
                                }}
                                placeholder={'选填'}
                                onChangeText={ this.updatePrice }
                                value={ this.state.Price }
                            />
                        </View>
                        <View style={{height: 10}} />
                        <View style={{flexDirection: 'row', backgroundColor: 'white'}}>
                            <Button
                                title='有效时间'
                                titleStyle={{color: 'steelblue', fontSize: 14, fontWeight: 'bold'}}
                                containerStyle={{flex: 1.1}}
                                buttonStyle={{backgroundColor: '#EFEFF5', flex: 1, alignItems: 'center'}}
                                raised={true}
                                onPress={() => {alert(this.state.validTime)}}
                            />
                            <SearchBar
                                searchIcon={false}
                                clearIcon={false}
                                cancelIcon={false}
                                lightTheme={true}
                                containerStyle={{
                                    flex: 4,
                                    height: 55,
                                    backgroundColor: 'white',
                                    borderWidth: 0, //no effect
                                    shadowColor: 'white', //no effect
                                    borderBottomColor: 'transparent',
                                    borderTopColor: 'transparent',
                                }}
                                inputContainerStyle={{
                                    marginTop: -9,
                                    height: 55,
                                    backgroundColor: 'white',
                                }}
                                placeholder={'必填。单位为天，信息有效期'}
                                onChangeText={ this.updateValidTime }
                                value={ this.state.validTime }
                            />
                        </View>
                        <View style={{height: 10}} />
                        <Button
                            title={'发布信息'}
                            titleStyle={{color: 'white', fontSize: 18}}
                            buttonStyle={{backgroundColor: 'red'}}
                            containerStyle={{width: 160, marginLeft: (width / 2 - 80)}}
                            raised={true}
                            onPress={() => {
                                if (this.state.validTime === '') {
                                    Alert.alert(
                                        '有效时间不可为空',
                                        '有效时间是必填项，它代表这条信息有效的时间跨度，单位为天',
                                        [
                                            {
                                                text: '好', onPress: () => {}
                                            }
                                        ],
                                        {cancelable: false},
                                    )
                                } else if (isValidTimeValid(this.state.validTime) === false) {
                                    Alert.alert(
                                        '有效时间格式不对',
                                        '有效时间是以天为单位的时间跨度，只能是正整数',
                                        [
                                            {
                                                text: '好', onPress: () => {}
                                            }
                                        ],
                                        {cancelable: false},
                                    )
                                } else if (this.state.Price === '' || isPriceValid(this.state.Price) === false) {
                                    Alert.alert(
                                        '价格格式不对',
                                        '价格只能是至多保留两位小数的正数',
                                        [
                                            {
                                                text: '好', onPress: () => {}
                                            }
                                        ],
                                        {cancelable: false},
                                    )
                                } else if (this.state.goodName === '') {
                                    Alert.alert(
                                        '物品名称不可为空',
                                        '物品名称是必填项，不可以空着哦',
                                        [
                                            {
                                                text: '好', onPress: () => {}
                                            }
                                        ],
                                        {cancelable: false},
                                    )
                                } else if (this.state.images.length === 0) {
                                    Alert.alert(
                                        '发布信息',
                                        ('您确定要发布该条' + buttons[this.state.selectedIndex] + '吗？'),
                                        [
                                            {
                                                text: '取消',
                                                //onPress: () => console.warn('Cancel Pressed'),
                                                style: 'cancel',
                                            },
                                            {
                                                text: '确定', onPress: () => {
                                                    let addType;
                                                    if (this.state.selectedIndex === 0)
                                                        addType = '/sellInfo';
                                                    else
                                                        addType = '/buyInfo';
                                                    let formData = new FormData();
                                                    formData.append('userID', Config.userInfo.userID);
                                                    formData.append('validTime', this.state.validTime);
                                                    formData.append('goodName', this.state.goodName);
                                                    if (this.state.description !== '')
                                                        formData.append('description', this.state.description);
                                                    if (this.state.Price !== '')
                                                        formData.append('price', this.state.Price);
                                                    HTTP.addInfo(addType, formData)
                                                        .then((response) => {
                                                            if (response.status === 1) {
                                                                Alert.alert(
                                                                    '发布成功',
                                                                    '成功发布该交易信息：' + this.state.goodName,
                                                                    [
                                                                        {
                                                                            text: '好', onPress: () => {
                                                                                this.keyID = 0;
                                                                                this.contentID = '';
                                                                                this.contentToken = '';
                                                                                this.uploadImageSuccess = false;
                                                                                this.setState({
                                                                                    goodName: '',
                                                                                    description: '',
                                                                                    Price: '',
                                                                                    validTime: '',
                                                                                    images: [],
                                                                                });
                                                                            }
                                                                        }
                                                                    ],
                                                                    {cancelable: false},
                                                                );
                                                            } else {
                                                                //console.warn(response);
                                                                Alert.alert(
                                                                    '出错啦',
                                                                    '网络可能出了问题，请再试一次吧',
                                                                    [
                                                                        {
                                                                            text: '好', onPress: () => {
                                                                            }
                                                                        }
                                                                    ],
                                                                    {cancelable: false},
                                                                )
                                                            }
                                                        })
                                                        .catch((error) => {
                                                            console.error(error);
                                                        });
                                                }
                                            }
                                        ]
                                    );
                                } else {
                                    Alert.alert(
                                        '发布信息',
                                        ('您确定要发布该条' + buttons[this.state.selectedIndex] + '吗？'),
                                        [
                                            {
                                                text: '取消',
                                                //onPress: () => console.warn('Cancel Pressed'),
                                                style: 'cancel',
                                            },
                                            {
                                                text: '确定', onPress: () => {
                                                    let params = {
                                                        path: this.state.images[0].uri,
                                                        type: 1,
                                                    };
                                                    //console.warn(params);
                                                    HTTP.addContent('/content', params)
                                                        .then((response) => {
                                                            if (response.status === 1) {
                                                                //console.warn('上传图片成功!');
                                                                //console.warn(response);
                                                                this.uploadImageSuccess = true;
                                                                this.contentID = response.contentID;
                                                                this.contentToken = response.contentToken;
                                                                let len = this.state.images.length;
                                                                for (let j = 1; this.uploadImageSuccess && j < len; ++j) {
                                                                    params = {
                                                                        path: this.state.images[j].uri,
                                                                        type: 1,
                                                                        contentID: this.contentID,
                                                                        contentToken: this.contentToken,
                                                                    };
                                                                    HTTP.addContent('/content', params)
                                                                        .then((response) => {
                                                                            if (response.status === 1) {
                                                                                //console.warn('上传图片成功!');
                                                                                //console.warn(response);
                                                                            } else {
                                                                                //console.warn(response);
                                                                                Alert.alert(
                                                                                    '出错啦',
                                                                                    '网络可能出了问题，请再试一次吧',
                                                                                    [
                                                                                        {
                                                                                            text: '好', onPress: () => {
                                                                                                this.uploadImageSuccess = false;
                                                                                            }
                                                                                        }
                                                                                    ],
                                                                                    {cancelable: false},
                                                                                )
                                                                            }
                                                                        }).catch((err) => {
                                                                        //console.warn('失败!');
                                                                        //console.warn(err);
                                                                    });
                                                                }
                                                                if (this.uploadImageSuccess) {
                                                                    let addType;
                                                                    if (this.state.selectedIndex === 0)
                                                                        addType = '/sellInfo';
                                                                    else
                                                                        addType = '/buyInfo';
                                                                    let formData = new FormData();
                                                                    formData.append('userID', Config.userInfo.userID);
                                                                    formData.append('validTime', this.state.validTime);
                                                                    formData.append('goodName', this.state.goodName);
                                                                    formData.append('contentID', this.contentID);
                                                                    formData.append('contentToken', this.contentToken);
                                                                    if (this.state.description !== '')
                                                                        formData.append('description', this.state.description);
                                                                    if (this.state.Price !== '')
                                                                        formData.append('price', this.state.Price);
                                                                    HTTP.addInfo(addType, formData)
                                                                        .then((response) => {
                                                                            if (response.status === 1) {
                                                                                Alert.alert(
                                                                                    '发布成功',
                                                                                    '成功发布该交易信息：' + this.state.goodName,
                                                                                    [
                                                                                        {
                                                                                            text: '好', onPress: () => {
                                                                                                this.keyID = 0;
                                                                                                this.contentID = '';
                                                                                                this.contentToken = '';
                                                                                                this.uploadImageSuccess = false;
                                                                                                this.setState({
                                                                                                    goodName: '',
                                                                                                    description: '',
                                                                                                    Price: '',
                                                                                                    validTime: '',
                                                                                                    images: [],
                                                                                                });
                                                                                            }
                                                                                        }
                                                                                    ],
                                                                                    {cancelable: false},
                                                                                );
                                                                            } else {
                                                                                //console.warn(response);
                                                                                Alert.alert(
                                                                                    '出错啦',
                                                                                    '网络可能出了问题，请再试一次吧',
                                                                                    [
                                                                                        {
                                                                                            text: '好', onPress: () => {
                                                                                            }
                                                                                        }
                                                                                    ],
                                                                                    {cancelable: false},
                                                                                )
                                                                            }
                                                                        })
                                                                        .catch((error) => {
                                                                            console.error(error);
                                                                        });
                                                                }
                                                            } else {
                                                                //console.warn(response);
                                                                Alert.alert(
                                                                    '出错啦',
                                                                    '网络可能出了问题，请再试一次吧',
                                                                    [
                                                                        {
                                                                            text: '好', onPress: () => {
                                                                            }
                                                                        }
                                                                    ],
                                                                    {cancelable: false},
                                                                )
                                                            }
                                                        }).catch((err) => {
                                                        //console.warn('失败!');
                                                        //console.warn(err);
                                                    });
                                                }
                                            },
                                        ],
                                        {cancelable: false},
                                    );
                                }
                            }}
                        />
                        <View style={{height: 10}} />
                    </View>
                </ScrollView>
            )
        }
    }
}
