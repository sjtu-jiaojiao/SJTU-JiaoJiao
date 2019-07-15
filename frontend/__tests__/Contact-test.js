import React from 'react';
import ContactScreen from '../src/Views/Contact/Contact';

import renderer from 'react-test-renderer';

test('renders correctly', () => {
    const tree = renderer.create(<ContactScreen/>).toJSON();
    expect(tree).toMatchSnapshot();
})
