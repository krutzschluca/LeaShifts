import { getStores, addStore, updateStore, deleteStore } from './services/api.js';

document.addEventListener('DOMContentLoaded', () => {
  const storeList = document.getElementById('storeList');
  const errorMessage = document.getElementById('errorMessage');
  const addStoreButton = document.getElementById('addStoreButton');
  const storeModal = new bootstrap.Modal(document.getElementById('storeModal'));
  const storeForm = document.getElementById('storeForm');
  const storeIdInput = document.getElementById('storeId');
  const addressInput = document.getElementById('address');
  const openingTimeInput = document.getElementById('openingTime');
  const closingTimeInput = document.getElementById('closingTime');

  const renderStores = (stores) => {
    storeList.innerHTML = '';
    stores.forEach(store => {
      const row = document.createElement('tr');
      row.innerHTML = `
        <td>${store.id}</td>
        <td>${store.address}</td>
        <td>${store.opening_time}</td>
        <td>${store.closing_time}</td>
        <td>
          <button class="btn btn-info btn-sm" data-store-id="${store.id}" data-action="edit">Edit</button>
          <button class="btn btn-danger btn-sm" data-store-id="${store.id}" data-action="delete">Delete</button>
        </td>
      `;
      storeList.appendChild(row);
    });

    document.querySelectorAll('button[data-action="edit"]').forEach(button => {
      button.addEventListener('click', (event) => {
        const storeId = event.target.getAttribute('data-store-id');
        const store = stores.find(store => store.id == storeId);
        if (store) {
          storeIdInput.value = store.id;
          addressInput.value = store.address;
          openingTimeInput.value = store.opening_time;
          closingTimeInput.value = store.closing_time;
          storeModal.show();
        }
      });
    });

    document.querySelectorAll('button[data-action="delete"]').forEach(button => {
      button.addEventListener('click', async (event) => {
        const storeId = event.target.getAttribute('data-store-id');
        try {
          await deleteStore(storeId);
          fetchAndRenderStores();
        } catch (error) {
          console.error('Error deleting store:', error);
          errorMessage.textContent = 'Failed to delete store. Maybe it has shifts assigned.';
        }
      });
    });
  };

  const fetchAndRenderStores = () => {
    getStores().then(stores => {
      renderStores(stores);
    }).catch(error => {
      console.error('Error fetching stores:', error);
      errorMessage.textContent = 'Failed to load stores. Please try again later.';
    });
  };

  addStoreButton.addEventListener('click', () => {
    storeForm.reset();
    storeIdInput.value = '';
    storeModal.show();
  });

  storeForm.addEventListener('submit', async (event) => {
    event.preventDefault();

    const storeData = {
      address: addressInput.value,
      opening_time: openingTimeInput.value,
      closing_time: closingTimeInput.value
    };

    try {
      if (storeIdInput.value) {
        await updateStore(storeIdInput.value, storeData);
      } else {
        await addStore(storeData);
      }
      storeModal.hide();
      fetchAndRenderStores();
    } catch (error) {
      console.error('Error saving store:', error);
      errorMessage.textContent = 'Failed to save store. Please try again later.';
    }
  });

  fetchAndRenderStores();
});
