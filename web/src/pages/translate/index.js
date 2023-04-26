import React, { useState } from 'react';

const Translate = () => {
  const [filePath, setFilePath] = useState('');
  const [sourceLang, setSourceLang] = useState('en');
  const [targetLang, setTargetLang] = useState('zh');
  const [translator, setTranslator] = useState('google');

  const handleFilePathChange = (e) => {
    setFilePath(e.target.value);
  };

  const handleSourceLangChange = (e) => {
    setSourceLang(e.target.value);
  };

  const handleTargetLangChange = (e) => {
    setTargetLang(e.target.value);
  };

  const handleTranslatorChange = (e) => {
    setTranslator(e.target.value);
  };

const handleTranslate = async () => {
    try {
      const formData = new FormData();
      formData.append('filePath', filePath);
      formData.append('sourceLang', sourceLang);
      formData.append('targetLang', targetLang);
      formData.append('translator', translator);

      const response = await fetch('/api/translate', {
        method: 'POST',
        body: formData,
      });
      const result = await response.json();
      console.log(result);
    } catch (error) {
      console.error(error);
    }
  };


  return (
    <div>
      <h1>翻译文件页面</h1>
      <div>
        <label htmlFor="filePath">文件路径：</label>
        <input type="text" id="filePath" value={filePath} onChange={handleFilePathChange} />
      </div>
      <div>
        <label htmlFor="sourceLang">源语言：</label>
        <select id="sourceLang" value={sourceLang} onChange={handleSourceLangChange}>
          <option value="en">英语</option>
          <option value="zh">中文</option>
        </select>
      </div>
      <div>
        <label htmlFor="targetLang">目标语言：</label>
        <select id="targetLang" value={targetLang} onChange={handleTargetLangChange}>
          <option value="zh">中文</option>
          <option value="en">英语</option>
        </select>
      </div>
      <div>
        <label htmlFor="translator">翻译器：</label>
        <select id="translator" value={translator} onChange={handleTranslatorChange}>
          <option value="google">Google</option>
        </select>
      </div>
      <button onClick={handleTranslate}>翻译</button>
    </div>
  );
};

export default Translate;
