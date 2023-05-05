import React, { useState } from 'react';
import {Upload, Translate} from '../index'

const steps = [
  { id: 1, title: '上传文件 =>' },
  { id: 2, title: '开始翻译文件 =>' },
  { id: 3, title: '下载文件' },
];

const Stepper = () => {
  const [activeStep, setActiveStep] = useState(0);

  const handleNext = () => {
    setActiveStep((prevActiveStep) => prevActiveStep + 1);
  };

  const handleBack = () => {
    setActiveStep((prevActiveStep) => prevActiveStep - 1);
  };

  return (
    <div style={{width: '80%', margin: 'auto', border: '2px solid black', padding: '20px'}}>
      <h1 style={{textAlign: 'center', backgroundColor: 'lightblue', padding: '10px'}}>epub-translator</h1>
      <div style={{display: 'flex', justifyContent: 'center'}}>
        {steps.map((step) => (
          <div key={step.id} style={{margin: '0 20px', border: '1px solid black', padding: '10px'}}>
            <span>{step.title}</span>
          </div>
        ))}
      </div>
      <div style={{border: '2px solid black', padding: '20px'}}>
        {activeStep === 0 && (
          <div>
            <Upload/>
            <button style={{margin: '10px', backgroundColor: 'lightblue', color: 'white', border: 'none', padding: '10px'}} onClick={handleNext}>下一步</button>
          </div>
        )}
        {activeStep === 1 && (
          <div>
            <Translate/>
            <button style={{margin: '10px', backgroundColor: 'lightblue', color: 'white', border: 'none', padding: '10px'}} onClick={handleBack}>上一步</button>
            <button style={{margin: '10px', backgroundColor: 'lightblue', color: 'white', border: 'none', padding: '10px'}} onClick={handleNext}>下一步</button>
          </div>
        )}
        {activeStep === 2 && (
          <div>
            <p style={{border: '1px solid black', padding: '10px'}}>下载文件组件</p>
            <button style={{margin: '10px', backgroundColor: 'lightblue', color: 'white', border: 'none', padding: '10px'}} onClick={handleBack}>上一步</button>
          </div>
        )}
      </div>
    </div>
  );
};

export default Stepper;
