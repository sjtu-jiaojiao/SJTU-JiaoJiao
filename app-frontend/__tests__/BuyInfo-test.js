import React from 'react';
import BuyInfoScreen from '../src/Views/BuyInfo/BuyInfo';

import renderer from 'react-test-renderer';

test('renders correctly', () => {
    const tree = renderer.create(<BuyInfoScreen />).toJSON();
    expect(tree).toMatchSnapshot();
});
