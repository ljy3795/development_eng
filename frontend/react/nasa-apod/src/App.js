import React, { Component } from 'react';
import './App.css';
import ViewerTemplate from './component/ViewerTemplate';
import SpaceNavigator from './component/SpaceNavigator';
import Viewer from './component/Viewer';

import moment from 'moment';
import * as api from './lib/api';

class App extends Component {

  // getAPOD = (date) => {
  //   // promise & then -> 비동기 방식
  //   // api.getAPOD(date)까지 먼저 실행하고 바로 다음으로 넘어감
  //   api.getAPOD(date).then((response) => {
  //     console.log(response)
  //   });
  // }

  // // async / await 방식 이용
  // getAPOD = async (date) => {
  //   try {
  //     const response = await api.getAPOD(date);
  //     console.log(response)
  //   } catch (e) {
  //     // error occur
  //     console.log(e);
  //   }
  // }

  state = {
    loading: false,
    maxDate: null,
    date: null,
    url: null,
    mediaType: null
  }

  // 1) 요청받으면 state:true
  // 2) api.getAPOD(date)로 data 받음
  // 3) 전달받은 값을 setState
  // 4) loading 종료 : state:false
  getAPOD = async (date) => {
    if (this.state.loading) return; // 이미 요청중이면 무시

    // 로딩 시작
    this.setState({
      loading: true
    });

    try {
      const response = await api.getAPOD(date);
      // async 할당 + 새로운 이름
      const { date: retrievedDate, url, media_type: mediaType } = response.data;

      if (!this.state.maxDate) {
        // maxDate가 없으면 지금 받은 date할당
        this.setState({
          maxDate: retrievedDate
        });
      }

      // 전달받은 데이터 넣어주기
      this.setState({
        date: retrievedDate,
        mediaType: mediaType,
        url: url
      });
    } catch (e) {
      console.log(e);
    }
      // loading 종료
      this.setState({
        loading:false
      })
  }
  
  handlePrev = () => {
    const { date } = this.state;
    const prevDate = moment(date).subtract(1,'days').format('YYYY-MM-DD');
    console.log(prevDate);
    this.getAPOD(prevDate);
  }

  handleNext = () => {
    const { date } = this.state;
    const nextDate = moment(date).add(1,'days').format('YYYY-MM-DD');
    console.log(nextDate);
    this.getAPOD(nextDate);
  }


  componentDidMount() {
    this.getAPOD();
  }

  render() {
    const {url, mediaType, loading } = this.state;

    return (
        <div className='App'>
          <ViewerTemplate
            spaceNaviagor={<SpaceNavigator 
                            onPrev={this.handlePrev}
                            onNext={this.handleNext}
                          />}
            // viewer={<Viewer
            //           url="https://www.youtube.com/embed/uj3Lq7Gu94Y?rel=0"
            //           mediaType="viedo"
            // />}
            viewer={<Viewer
                      url={url}
                      mediaType={mediaType}
                      loading={loading}
                  />}
          />
        </div>
    );
  }
}

export default App;
