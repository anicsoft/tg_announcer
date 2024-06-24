import { ActionIcon, BoxComponentProps, Button, Checkbox, CloseButton, Divider, FileInput, Flex, Group, Input, MantineStyleProp, Modal, MultiSelect, Select, Stack, Switch, Text, TextInput, Title, rem } from '@mantine/core'
import { IconAt, IconClock, IconPhoto } from '@tabler/icons-react'
import React, { useEffect, useMemo, useRef, useState } from 'react'
import RichTexInput from '../ui/RichTexInput';
import { useForm } from '@mantine/form';
import { useQuery } from '@tanstack/react-query';
import { DatePickerInput, TimeInput } from '@mantine/dates';
import { useDisclosure } from '@mantine/hooks';
import OfferCard from '../components/OfferCard';
import { Offer, OfferRequest } from '../utils/data';
import OfferModal from '../ui/OfferModal';

export default function AdminConsole() {


  const [opened, { open, close }] = useDisclosure(false);
  const [hasPromocode, setHasPromocode] = useState(false)
  const [singleDayOffer, setSingleDayOffer] = useState(true)
  const [dateRange, setDateRange] = useState([])
  const [date, setDate] = useState()
  const [image, setImage] = useState<File | undefined | null>()
  const [content, setContent] = useState()
  const [innerDate, setInnerDate] = useState(Date.now())
  const [startTime, setStartTime] = useState()
  const [endTime, setEndTime] = useState()
  const [newOffer, setNewOffer] = useState<OfferRequest | undefined>()
  const startTimeRef = useRef<HTMLInputElement>(null);
  const endTimeRef = useRef<HTMLInputElement>(null);


  const stackProps = {
    align: "stretch",
    justify: "center",
    gap: "md",
    w: "100%",
    my: "sm",
  }

  const inputLabelStyles: MantineStyleProp = { textAlign: "left" }


  const form = useForm({
    mode: 'uncontrolled',
    initialValues: {
      title: '',
      hasPromocode: hasPromocode,
      promocode: '',
      content: '',
      dateRange: [],
      date: null,
      startTime: '',
      endTime: '',
      offerCategory: [],
      offerImage: image,
      singleDayOffer: true, // DK if it is needed in the form
    },

    validate: {
      // email: (value) => (/^\S+@\S+$/.test(value) ? null : 'Invalid email'),
    },
    onValuesChange: (values) => {
      const offer = {
        company_id: 1,
        title: values.title,
        image: values.offerImage,
        categories: values.offerCategory,
        start_date_time: values.dateRange[0],
        end_date_time: values.dateRange[1],
        promo_code: values.promocode
      } as OfferRequest
      console.log('NEW OFFER', offer);

      setNewOffer(offer)
    },
  });


  const handlePromocodeSwitchChange = (isChecked) => {
    form.setValues({
      hasPromocode: isChecked
    })
    setHasPromocode(isChecked)
  }

  const handleSingleDayOfferSwitchChange = (isChecked) => {
    form.setValues({
      singleDayOffer: isChecked
    })
    setSingleDayOffer(isChecked)
  }

  const handleDateChange = (value) => {
    console.log(value);

    form.setValues({
      date: value
    })
    setDateRange([new Date(value), new Date(value)])
    setDate(value)
  }

  useEffect(() => {

    console.log("CHNGING DATE RANGE");

    const existingDate = dateRange;
    const startDate = new Date(existingDate[0]);
    // Set the time components
    startDate.setHours(0);
    startDate.setMinutes(0);
    const endDate = new Date(existingDate[0]);
    endDate.setHours(23);
    endDate.setMinutes(59);

    const newDateRange = [startDate, endDate]
    form.setValues({
      dateRange: newDateRange
    })

  }, [dateRange, singleDayOffer])

  const handleTimeChange = (value: string, start: boolean) => {
    // console.log("handle time change");
    // console.log(value);

    const timeParts = value.split(':');
    const hours = parseInt(timeParts[0], 10);
    const minutes = parseInt(timeParts[1], 10);

    // Create a new Date object with the existing date
    // const newDate = new Date(innerDate)
    // console.log(date);

    const existingDate = dateRange;
    if (start) {
      setStartTime(value)
      form.setValues({
        startTime: value
      })

      const newDate = new Date(existingDate[0]);

      // Set the time components
      newDate.setHours(hours);
      newDate.setMinutes(minutes);

      const newDateRange = [...dateRange]
      newDateRange.splice(0, 1, newDate)
      form.setValues({
        dateRange: newDateRange
      })
      setDateRange(newDateRange)
    } else {

      setEndTime(value)

      form.setValues({
        endTime: value
      })


      const newDate = new Date(existingDate[1]);
      // Set the time components
      newDate.setHours(hours);
      newDate.setMinutes(minutes);

      const newDateRange = [...dateRange]
      newDateRange.splice(1, 1, newDate)
      form.setValues({
        dateRange: newDateRange
      })
      setDateRange(newDateRange)

    }
  }

  const { data: offerCategories } = useQuery({
    queryKey: ['offerCategories'],
    queryFn: () => fetch('http://0.0.0.0:8888/categories/offer').then((res) =>
      res.json(),
    )
  })

  const startTimePickerControl = (
    <ActionIcon variant="subtle" color="gray" onClick={() => startTimeRef.current?.showPicker()}>
      <IconClock style={{ width: rem(16), height: rem(16) }} stroke={1.5} />
    </ActionIcon>
  );
  const endTimePickerControl = (
    <ActionIcon variant="subtle" color="gray" onClick={() => endTimeRef.current?.showPicker()}>
      <IconClock style={{ width: rem(16), height: rem(16) }} stroke={1.5} />
    </ActionIcon>
  );


  const getTimeSlots = (startTime = "00:00", endTime = "24:00") => {
    const start = 0; // Start time in hours
    const end = 24; // End time in hours

    const slots = Array.from({ length: (end - start) * 2 }, (_, index) => {
      const hour = start + Math.floor(index / 2);
      const minute = index % 2 === 0 ? '00' : '30';
      return `${String(hour).padStart(2, '0')}:${minute}`;
    });

    const filteredSlots = slots.filter(slot => {
      return (startTime === "00:00" || slot > startTime) && (endTime === "24:00" || slot < endTime);
    });
    // console.log(filteredSlots);

    return filteredSlots
  }

  const StartTimeList = () => {
    const slots = getTimeSlots(undefined, endTime)

    return (
      <>
        <datalist id="startTime">
          {slots.map((slot) =>
            <option key={slot} value={slot} >{slot}</option>
          )}
        </datalist>
      </>
    )
  }
  const EndTimeList = () => {
    const slots = getTimeSlots(startTime, undefined)
    return (
      <>
        <datalist id="endTime">
          {slots.map((slot) =>
            <option key={slot} value={slot} >{slot}</option>
          )}
        </datalist>
      </>
    )
  }

  const handleOfferContentChange = (value) => {
    form.setValues({
      content: value
    })
  }


  const CheckboxComp = useMemo(() => {
    return (
      <Switch
        style={inputLabelStyles}
        label="Offer has a promocode"
        key={form.key('hasPromocode')}
        flex={1}
        type='checkbox'
        checked={hasPromocode}
        onChange={(event) => handlePromocodeSwitchChange(event.currentTarget.checked)}
      />
      // <Checkbox
      //   checked={hasPromocode}
      //   style={inputLabelStyles}
      //   label="Offer has a promocode"
      //   key={form.key('hasPromocode')}
      //   onChange={(event) => handlePromocodeSwitchChange(event.currentTarget.checked)}
      // />
    )
  }, [hasPromocode])

  // const RichTextEditorComp = useMemo(() => {
  //   return (
  //     <RichTexInput initcontent={content}></RichTexInput>
  //   )
  // }, [content])

  const handleImageUpload = async (value) => {
    console.log(value);
    console.log(await value.text());
    console.log(await value.arrayBuffer());
    const reader = new FileReader();
    const blob = reader.readAsDataURL(value)
    console.log(blob);

    setImage(value)
  }

  return (
    <>
      <Title order={1} size={"h2"} py="md">Add new offer</Title>
      <form onSubmit={form.onSubmit((values) => { console.log(values) })}>

        <Flex
          align={"flex-start"}
          direction={"column"}
          px="lg"
          pb="lg"
          gap={48}

        >
          <Stack
            {...stackProps}
          >
            <TextInput
              required
              label="Title"
              style={inputLabelStyles}
              key={form.key('title')}
              flex={1}
              {...form.getInputProps('title')}
            // placeholder="Title"
            // description="Description below the input"
            // inputWrapperOrder={['label', 'error', 'input', 'description']}
            />
            <MultiSelect
              style={inputLabelStyles}
              label="Offer type"
              placeholder="Pick value"
              required
              size="md"
              key={form.key('offerCategory')}
              data={offerCategories ? offerCategories.map(cat => cat.name) : []}
              {...form.getInputProps('offerCategory')}
            />
            <FileInput
              style={inputLabelStyles}
              required
              clearable
              rightSection={<IconPhoto style={{ width: rem(18), height: rem(18) }} stroke={1.5} />}
              accept="image/png,image/jpeg"
              label="Upload offer picture"
              placeholder="Upload picture"
              key={form.key('offerImage')}
              // value={form.getValues().offerImage}
              // onChange={(value) => handleImageUpload(value)}
              // valueComponent={() => {
              //   const url = URL.createObjectURL(new Blob(await form.getValues().offerImage?.arrayBuffer))
              //   return <Image
              //   radius="md"
              //   src={{url}}
              // />
              // }}
              {...form.getInputProps('offerImage')}
            />
          </Stack>
          <Stack
            {...stackProps}
          >

            <Title ta="left" order={3}>Promocode</Title>
            {/* <Switch
              style={inputLabelStyles}
              label="Offer has a promocode"
              key={form.key('hasPromocode')}
              flex={1}
              type='checkbox'
              checked={hasPromocode}
              onChange={(event) => handlePromocodeSwitchChange(event.currentTarget.checked)}
            /> */}
            <Checkbox
              checked={hasPromocode}
              style={inputLabelStyles}
              label="Offer has a promocode"
              key={form.key('hasPromocode')}
              onChange={(event) => handlePromocodeSwitchChange(event.currentTarget.checked)}
            />
            {/* <CheckboxComp></CheckboxComp> */}
            {/* {CheckboxComp} */}
            <TextInput
              disabled={!hasPromocode}
              flex={1}
              style={inputLabelStyles}
              label="Promocode"
              key={form.key('promocode')}
              {...form.getInputProps('promocode')}
            />


          </Stack>
          <Stack
            {...stackProps}
          >

            <Title ta="left" order={3}>Offer content</Title>
            <Input.Wrapper
              required

              style={inputLabelStyles}
              // label="Some label"
              // description="Some description"
              label="Offer content"
              // placeholder="Custom layout"
              description="Write offer content here"

              key={form.key('content')}
              {...form.getInputProps('content')}
            // error="Some error"
            >
              {/* <RichTexInput content={form.getValues().content} onChange={handleOfferContentChange}></RichTexInput> */}
              <RichTexInput initcontent={form.getValues().content}></RichTexInput>

              {/* {RichTextEditorComp} */}
              {/* <Input<any> component={TextEditor} /> */}
            </Input.Wrapper>
          </Stack>
          {/* <Divider my="md" h={1} w={"100%"} /> */}
          <Stack
            {...stackProps}
          >
            <Title order={2} size={"h3"} ta="left">Offer duration</Title>
            <DatePickerInput
              style={inputLabelStyles}
              label="Pick date"
              placeholder="Pick single date"
              value={date}
              onChange={(value) => handleDateChange(value)}
            />
            {/* <Switch
              style={inputLabelStyles}
              label="Whole day"
              key={form.key('singleDayOffer')}
              defaultChecked={singleDayOffer}
              type='checkbox'
              checked={singleDayOffer}
              onChange={(event) => handleSingleDayOfferSwitchChange(event.currentTarget.checked)}
            /> */}
            <Checkbox
              required
              disabled={!form.getValues().date}
              style={inputLabelStyles}
              label="Whole day"
              key={form.key('singleDayOffer')}
              defaultChecked={singleDayOffer}
              type='checkbox'
              checked={singleDayOffer}
              onChange={(event) => handleSingleDayOfferSwitchChange(event.currentTarget.checked)}
            />

            <Group
              flex={1}
            >

              <TimeInput
                style={inputLabelStyles}
                label="Start time"
                withSeconds={false}
                flex={1}
                list="startTime"
                disabled={!form.getValues().date || singleDayOffer}
                maxTime={endTime}
                placeholder="Start time"
                ref={startTimeRef}
                rightSection={startTimePickerControl}
                value={form.getValues().startTime}
                onChange={(event) => handleTimeChange(event.currentTarget.value, true)}
              />
              <TimeInput
                style={inputLabelStyles}
                label="End time"
                list="endTime"
                withSeconds={false}
                minTime={startTime}
                disabled={!form.getValues().date || singleDayOffer}
                flex={1}
                placeholder="End time"
                ref={endTimeRef}
                rightSection={endTimePickerControl}
                value={form.getValues().endTime}
                onChange={(event) => handleTimeChange(event.currentTarget.value, false)}
              />
              <StartTimeList></StartTimeList>
              <EndTimeList></EndTimeList>
            </Group>

          </Stack>
          <Group w="100%">
            <Button flex={1} type="submit">Submit</Button>
            <Button flex={1} type="button" variant='light' onClick={open} disabled={newOffer == undefined}>Preview</Button>
          </Group>
        </Flex>
      </form>
      {newOffer &&
        <OfferModal opened={opened} onClose={close} offer={newOffer}></OfferModal>
        // <Modal opened={opened} onClose={close} title={newOffer?.title} padding={0} styles={{
        //   header: { position: "absolute", width: "100%", backgroundColor: "transparent" },
        //   // content: { paddingInline: "20px" }
        // }}>
        //   <OfferCard popUp={newOffer}></OfferCard>
        // </Modal>
      }
    </>
  )
}
