export async function fetchStocks() {
    const response = await fetch("http://localhost:8080/stocks");
    
    if (!response.ok) {
      throw new Error(`Server error: ${response.statusText}`);
    }
    
    
    return await response.json();
  }