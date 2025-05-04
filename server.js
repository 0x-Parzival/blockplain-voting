const express = require('express');
const fs = require('fs');
const path = require('path');
const app = express();
const PORT = 3000;

// Serve static files like index.html and CSS
app.use(express.static(__dirname));
app.use(express.json());

// Serve index.html explicitly
app.get('/', (req, res) => {
  res.sendFile(path.join(__dirname, 'index.html'));
});

// Ensure data directory exists
const dataDir = path.join(__dirname, 'data');
if (!fs.existsSync(dataDir)) {
  fs.mkdirSync(dataDir);
}

// Handle block save
app.post('/save-block', (req, res) => {
  const { vote, regionName } = req.body;

  const colorMap = {
    A: '#ff5733',
    B: '#33aaff'
  };

  const hexCode = colorMap[vote] || '#000000';
  const line = `Region: ${regionName}, Vote: ${vote}, Hex: ${hexCode}\n`;

  fs.appendFile(path.join(dataDir, 'blocks.txt'), line, err => {
    if (err) {
      console.error('Error writing to file:', err);
      return res.status(500).send('Error saving data.');
    }
    res.send('Saved successfully.');
  });
});

app.listen(PORT, () => {
  console.log(`Server is running at http://localhost:${PORT}`);
});
