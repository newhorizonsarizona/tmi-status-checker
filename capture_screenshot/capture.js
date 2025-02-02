const puppeteer = require("puppeteer");

(async () => {
  // Launch a headless browser
  const browser = await puppeteer.launch({ args: ['--no-sandbox', '--disable-setuid-sandbox'] });
  const page = await browser.newPage();

  // Navigate to the webpage
  await page.goto(
    "https://dashboards.toastmasters.org/ClubReport.aspx?id="+process.env.CLUB_NUMBER,
  );

  // Wait for the element with the class name to load (modify the class selector as needed)
  const selector = ".tabBody"; // Change this to your actual CSS class or selector

  await page.waitForSelector(selector);

  // Get the element and take a screenshot of it
  const element = await page.$(selector);
  await element.screenshot({ path: "../reports/dcp_report.png" });

  // Close the browser
  await browser.close();

  console.log("Screenshot saved as ../reports/dcp_report.png");
})();
