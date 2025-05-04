const express = require('express');
const fs = require('fs');
const path = require('path');
const app = express();

app.use(express.json());
app.use(express.static('public'));

app.post('/save-block', (req, res) => {
  const { vote, regionName } = req.body;
  const colorMap = {
    A: '#ff5733',
    B: '#33aaff'
  };
  const hexCode = colorMap[vote] || '#000000';

  const line = `Region: ${regionName}, Vote: ${vote}, Hex: ${hexCode}\n`;

  const filePath = path.join(__dirname, 'data', 'blocks.txt');
  fs.appendFile(filePath, line, err => {
    if (err) {
      console.error('Write error:', err);
      return res.status(500).send('Error saving block.');
    }
    res.send('Block saved successfully.');
  });
});

app.listen(3000, () => {
  console.log('Server running at http://localhost:3000');
});
